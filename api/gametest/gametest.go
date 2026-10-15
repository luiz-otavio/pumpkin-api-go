package gametest

import (
	"errors"

	bindings "github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_gametest"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/imports/pumpkin_plugin_world"
	witTypes "go.bytecodealliance.org/pkg/wit/types"
)

type Test = bindings.Test
type RegistrationBuilder = bindings.RegistrationBuilder
type BlockState = pumpkin_plugin_world.BlockState
type Entity = pumpkin_plugin_world.Entity

type callback interface {
	~func() | ~func(*Test) | ~func(BlockState) bool | ~func(*Entity) bool
}

func addCallback[Handler callback](callbacks *registry[Handler], handler Handler) uint32 {
	if handler == nil {
		panic("gametest: nil callback")
	}
	return callbacks.add(handler)
}

var (
	voidCallbacks   registry[func()]
	testCallbacks   registry[func(*Test)]
	asyncCallbacks  registry[func(*Test)]
	blockCallbacks  registry[func(BlockState) bool]
	entityCallbacks registry[func(*Entity) bool]
)

func Register(className, testName string, handler func(*Test)) (*RegistrationBuilder, error) {
	id := addCallback(&testCallbacks, handler)
	result := bindings.Register(className, testName, bindings.TestCallbackId{Id: id})
	if result.Tag() == witTypes.ResultErr {
		testCallbacks.remove(id)
		return nil, errors.New(result.Err())
	}
	return result.Ok(), nil
}

func RegisterAsync(className, testName string, handler func(*Test)) (*RegistrationBuilder, error) {
	id := addCallback(&asyncCallbacks, handler)
	result := bindings.RegisterAsync(className, testName, bindings.AsyncTestCallbackId{Id: id})
	if result.Tag() == witTypes.ResultErr {
		asyncCallbacks.remove(id)
		return nil, errors.New(result.Err())
	}
	return result.Ok(), nil
}

func RegisterVoidCallback(handler func()) bindings.VoidCallbackId {
	return bindings.VoidCallbackId{Id: addCallback(&voidCallbacks, handler)}
}

func RegisterBlockPredicate(handler func(BlockState) bool) bindings.BlockPredicateCallbackId {
	return bindings.BlockPredicateCallbackId{Id: addCallback(&blockCallbacks, handler)}
}

func RegisterEntityPredicate(handler func(*Entity) bool) bindings.EntityPredicateCallbackId {
	return bindings.EntityPredicateCallbackId{Id: addCallback(&entityCallbacks, handler)}
}

func UnregisterVoidCallback(id bindings.VoidCallbackId) {
	voidCallbacks.remove(id.Id)
}

func UnregisterBlockPredicate(id bindings.BlockPredicateCallbackId) {
	blockCallbacks.remove(id.Id)
}

func UnregisterEntityPredicate(id bindings.EntityPredicateCallbackId) {
	entityCallbacks.remove(id.Id)
}

func InvokeVoid(id bindings.VoidCallbackId) {
	voidCallbacks.get(id.Id)()
}

func InvokeTest(id bindings.TestCallbackId, test *Test) {
	testCallbacks.get(id.Id)(test)
}

func InvokeAsyncTest(id bindings.AsyncTestCallbackId, test *Test) {
	asyncCallbacks.get(id.Id)(test)
}

func InvokeBlockPredicate(id bindings.BlockPredicateCallbackId, block BlockState) bool {
	return blockCallbacks.get(id.Id)(block)
}

func InvokeEntityPredicate(id bindings.EntityPredicateCallbackId, entity *Entity) bool {
	return entityCallbacks.get(id.Id)(entity)
}
