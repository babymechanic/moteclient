package components

import (
	"github.com/babymechanic/client/utils"
	"github.com/babymechanic/common/messages"
)

type Display struct {
	invoker *utils.RPCInvoker
}

func NewDisplay(invoker *utils.RPCInvoker) *Display {
	return &Display{invoker}
}

func (display *Display) Resolution() messages.Resolution {
	var response messages.Resolution
	display.invoker.Invoke("Display.Resolution", struct{}{}, &response)
	return response
}

func init() {

}
