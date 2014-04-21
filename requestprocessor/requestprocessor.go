package requestprocessor

import (
	"fmt"
	"github.com/babymechanic/motecommon/messages"
)

type RequestProcessor struct {
	listenPort      int
	commandExecutor CommandExecutor
}

func New(listenPort int) *RequestProcessor {
	return &RequestProcessor{listenPort: listenPort}
}

func (requestProcessor *RequestProcessor) ProcessRequests() {
	returnValue, _ := requestProcessor.commandExecutor.Execute("Display.Resolution", nil).(messages.Resolution)
	fmt.Println(returnValue.Width, " x ", returnValue.Height)
	requestProcessor.commandExecutor.Execute("Display.SetResolution", messages.Resolution{Width: 1920, Height: 1080})
}
