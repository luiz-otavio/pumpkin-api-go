package api

import (
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_command"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_context"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_event"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_metadata"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_server"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_world"
)

type Metadata = pumpkin_plugin_metadata.PluginMetadata

type Plugin interface {
	Metadata() Metadata
	OnLoad(ctx *pumpkin_plugin_context.Context)
	OnUnload(ctx *pumpkin_plugin_context.Context)
	HandleEvent(eventId uint32, server *pumpkin_plugin_server.Server, event pumpkin_plugin_event.Event) pumpkin_plugin_event.Event
	HandleCommand(commandId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, args *pumpkin_plugin_command.ConsumedArgs) int32
	HandleCommandSuggestion(handlerId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, request pumpkin_plugin_command.SuggestionRequest) pumpkin_plugin_command.CommandSuggestions
	HandleTask(handlerId uint32, server *pumpkin_plugin_server.Server)
	HandleIpcMessage(sender string, message []uint8) []uint8
	HandleAiGoalCanStart(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) bool
	HandleAiGoalShouldContinue(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) bool
	HandleAiGoalStart(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity)
	HandleAiGoalTick(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity)
	HandleAiGoalStop(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity)
	HandleGeneratePhase(generatorId uint32, phase pumpkin_plugin_world.GenerationPhase, chunk *pumpkin_plugin_world.ChunkBuffer)
}

var instance Plugin

func RegisterPlugin(p Plugin) {
	instance = p
}

func GetInstance() Plugin {
	return instance
}

type DefaultPlugin struct{}

func (p *DefaultPlugin) Metadata() Metadata {
	return Metadata{}
}

func (p *DefaultPlugin) OnLoad(ctx *pumpkin_plugin_context.Context)   {}
func (p *DefaultPlugin) OnUnload(ctx *pumpkin_plugin_context.Context) {}

func (p *DefaultPlugin) HandleEvent(eventId uint32, server *pumpkin_plugin_server.Server, event pumpkin_plugin_event.Event) pumpkin_plugin_event.Event {
	return event
}

func (p *DefaultPlugin) HandleCommand(commandId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, args *pumpkin_plugin_command.ConsumedArgs) int32 {
	return 0
}

func (p *DefaultPlugin) HandleCommandSuggestion(handlerId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, request pumpkin_plugin_command.SuggestionRequest) pumpkin_plugin_command.CommandSuggestions {
	return pumpkin_plugin_command.CommandSuggestions{}
}

func (p *DefaultPlugin) HandleTask(handlerId uint32, server *pumpkin_plugin_server.Server) {}

func (p *DefaultPlugin) HandleIpcMessage(sender string, message []uint8) []uint8 {
	return nil
}

func (p *DefaultPlugin) HandleAiGoalCanStart(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) bool {
	return false
}

func (p *DefaultPlugin) HandleAiGoalShouldContinue(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) bool {
	return false
}

func (p *DefaultPlugin) HandleAiGoalStart(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) {
}

func (p *DefaultPlugin) HandleAiGoalTick(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) {
}

func (p *DefaultPlugin) HandleAiGoalStop(goalId uint32, server *pumpkin_plugin_server.Server, entity *pumpkin_plugin_world.Entity) {
}

func (p *DefaultPlugin) HandleGeneratePhase(generatorId uint32, phase pumpkin_plugin_world.GenerationPhase, chunk *pumpkin_plugin_world.ChunkBuffer) {
}
