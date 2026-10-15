package gametest

import (
	"testing"

	bindings "github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_gametest"
)

func TestCallbacks(t *testing.T) {
	called := 0
	voidID := RegisterVoidCallback(func() { called++ })
	defer UnregisterVoidCallback(voidID)
	InvokeVoid(voidID)
	InvokeVoid(voidID)
	if called != 2 {
		t.Fatalf("callback called %d times, want 2", called)
	}

	resource := new(Test)
	for _, asynchronous := range []bool{false, true} {
		handlers := &testCallbacks
		invoke := InvokeTest
		if asynchronous {
			handlers = &asyncCallbacks
			invoke = func(id bindings.TestCallbackId, test *Test) {
				InvokeAsyncTest(bindings.AsyncTestCallbackId{Id: id.Id}, test)
			}
		}
		var received *Test
		id := handlers.add(func(test *Test) { received = test })
		invoke(bindings.TestCallbackId{Id: id}, resource)
		handlers.remove(id)
		if received != resource {
			t.Fatalf("async=%v: callback did not receive the test resource", asynchronous)
		}
	}

	block := BlockState{}
	blockID := RegisterBlockPredicate(func(received BlockState) bool { return true })
	defer UnregisterBlockPredicate(blockID)
	if !InvokeBlockPredicate(blockID, block) {
		t.Fatal("block predicate result was not forwarded")
	}

	entity := new(Entity)
	entityID := RegisterEntityPredicate(func(received *Entity) bool { return received == entity })
	defer UnregisterEntityPredicate(entityID)
	if !InvokeEntityPredicate(entityID, entity) || InvokeEntityPredicate(entityID, nil) {
		t.Fatal("entity predicate argument or result was not forwarded")
	}
}

func TestNilCallback(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for a nil callback")
		}
	}()
	RegisterVoidCallback(nil)
}

func TestCallbackCanRegisterCallback(t *testing.T) {
	id := RegisterVoidCallback(func() {
		nested := RegisterVoidCallback(func() {})
		InvokeVoid(nested)
		UnregisterVoidCallback(nested)
	})
	defer UnregisterVoidCallback(id)
	InvokeVoid(id)
}
