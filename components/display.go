package components

import (
	"github.com/babymechanic/moteclient/utils"
	"github.com/babymechanic/motecommon/messages"
	"github.com/mitchellh/mapstructure"
	"net/rpc"
	"reflect"
)

type Display struct {
	invoker *utils.RPCInvoker
}

func (display *Display) Initialize(client *rpc.Client) {
	display.invoker = utils.NewRPCInvoker(client)
}

func (display *Display) Resolution() messages.Resolution {
	var response messages.Resolution
	display.invoker.Invoke("Display.Resolution", struct{}{}, &response)
	return response
}

func (display *Display) SetResolution(resolution interface{}) {
	castedResolution := messages.Resolution{}
	mapstructure.Decode(resolution, &castedResolution)
	var response messages.Resolution
	display.invoker.Invoke("Display.SetResolution", castedResolution, &response)
}

func init() {
	display := &Display{}
	Register("Display.Resolution", display, reflect.TypeOf(nil), reflect.TypeOf(messages.Resolution{}))
	Register("Display.SetResolution", display, reflect.TypeOf(messages.Resolution{}), reflect.TypeOf(nil))
}
