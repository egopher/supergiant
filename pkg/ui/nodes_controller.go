package ui

import (
	"net/http"

	"github.com/supergiant/supergiant/pkg/client"
	"github.com/supergiant/supergiant/pkg/models"
)

func NewNode(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, "nodes/new.html", map[string]interface{}{
		"title":      "Nodes",
		"formAction": "/ui/nodes",
		"model": map[string]interface{}{
			"kube_id": nil,
			"class":   "",
		},
	})
}

func CreateNode(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	m := new(models.Node)
	if err := unmarshalFormInto(r, m); err != nil {
		return err
	}
	if err := sg.Nodes.Create(m); err != nil {
		return renderTemplate(w, "nodes/new.html", map[string]interface{}{
			"title":      "Nodes",
			"formAction": "/ui/nodes",
			"model":      m,
			"error":      err.Error(),
		})
	}

	http.Redirect(w, r, "/ui/nodes", http.StatusFound)
	return nil
}

func ListNodes(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	fields := []map[string]interface{}{
		{
			"title": "Kube ID",
			"type":  "field_value",
			"field": "kube_id",
		},
		{
			"title": "Class",
			"type":  "field_value",
			"field": "class",
		},
		{
			"title": "Provider ID",
			"type":  "field_value",
			"field": "provider_id",
		},
		{
			"title":             "CPU Usage",
			"type":              "percentage",
			"numerator_field":   "cpu_usage",
			"denominator_field": "cpu_limit",
		},
		{
			"title":             "RAM usage",
			"type":              "percentage",
			"numerator_field":   "ram_usage",
			"denominator_field": "ram_limit",
		},
	}
	return renderTemplate(w, "nodes/index.html", map[string]interface{}{
		"title":       "Nodes",
		"uiBasePath":  "/ui/nodes",
		"apiListPath": "/api/v0/nodes",
		"fields":      fields,
		"showNewLink": true,
		"batchActionPaths": map[string]string{
			"Delete": "/delete",
		},
	})
}

func GetNode(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return err
	}
	item := new(models.Node)
	if err := sg.Nodes.Get(id, item); err != nil {
		return err
	}
	return renderTemplate(w, "nodes/show.html", map[string]interface{}{
		"title": "Nodes",
		"model": item,
	})
}

func DeleteNode(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return err
	}
	item := new(models.Node)
	item.ID = id
	if err := sg.Nodes.Delete(item); err != nil {
		return err
	}
	// http.Redirect(w, r, "/ui/nodes", http.StatusFound)
	return nil
}
