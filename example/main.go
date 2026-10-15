package main

import (
	"github.com/Pumpkin-MC/pumpkin-api-go/api"
	_ "github.com/Pumpkin-MC/pumpkin-api-go/pkg/exports/wit_exports"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_context"
)

type MyPlugin struct {
	api.DefaultPlugin
}

func (p *MyPlugin) Metadata() api.Metadata {
	return api.Metadata{
		Name:    "GoExamplePlugin",
		Version: "1.0.0",
		Authors: []string{"Alex"},
	}
}

func (p *MyPlugin) OnLoad(ctx *pumpkin_plugin_context.Context) {
}

func init() {
	api.RegisterPlugin(&MyPlugin{})
}

func main() {}
