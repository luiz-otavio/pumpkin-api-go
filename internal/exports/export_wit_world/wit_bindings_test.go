package export_wit_world

import (
	"testing"

	"github.com/Pumpkin-MC/pumpkin-api-go/api"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_server"
)

type taskPlugin struct {
	api.DefaultPlugin
	called    bool
	handlerID uint32
}

func (p *taskPlugin) Metadata() api.Metadata {
	return api.Metadata{}
}

func (p *taskPlugin) HandleTask(handlerID uint32, server *pumpkin_plugin_server.Server) {
	p.called = true
	p.handlerID = handlerID
}

func TestHandleTaskDispatchesToPlugin(t *testing.T) {
	plugin := &taskPlugin{}
	api.RegisterPlugin(plugin)

	HandleTask(42, nil)

	if !plugin.called {
		t.Fatal("HandleTask did not invoke the plugin callback")
	}
	if plugin.handlerID != 42 {
		t.Fatalf("HandleTask passed handler ID %d, want 42", plugin.handlerID)
	}
}

var _ api.Plugin = (*taskPlugin)(nil)
