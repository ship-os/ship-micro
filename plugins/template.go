package plugin

var (
	tmpl = `
package main

import (
	"{{.Path}}"
	"github.com/ship-os/ship-micro/plugins"
)

var Plugin = plugin.Plugin{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
