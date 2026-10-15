package main

import (
	"github.com/Pumpkin-MC/pumpkin-api-go/api"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_context"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_logging"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_scheduler"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_server"
	_ "github.com/Pumpkin-MC/pumpkin-api-go/pkg/wit_exports"
)

type MyPlugin struct {
	api.DefaultPlugin
	taskID uint32
}

func (p *MyPlugin) Metadata() api.Metadata {
	return api.Metadata{
		Name:    "GoExamplePlugin",
		Version: "1.0.0",
		Authors: []string{"Alex"},
	}
}

func (p *MyPlugin) OnLoad(ctx *pumpkin_plugin_context.Context) {
	p.taskID = pumpkin_plugin_scheduler.ScheduleRepeatingTask(1, 20, 20)
}

func (p *MyPlugin) OnUnload(ctx *pumpkin_plugin_context.Context) {
	if p.taskID != 0 {
		pumpkin_plugin_scheduler.CancelTask(p.taskID)
		p.taskID = 0
	}
}

func (p *MyPlugin) HandleTask(handlerId uint32, server *pumpkin_plugin_server.Server) {
	if handlerId == 1 {
		pumpkin_plugin_logging.Log(pumpkin_plugin_logging.LevelInfo, "Scheduled task fired")
	}
}

func init() {
	api.RegisterPlugin(&MyPlugin{})
}

func main() {}
