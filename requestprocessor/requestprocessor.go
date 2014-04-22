package requestprocessor

import (
	"fmt"
	"github.com/babymechanic/moteclient/requestprocessor/handlers"
	"net/http"
)

func New(listenPort int) *RequestProcessor {
	return &RequestProcessor{listenPort: listenPort}
}

type RequestProcessor struct {
	listenPort int
}

func (requestProcessor *RequestProcessor) ProcessRequests() {
	for route, handler := range handlers.Handlers() {
		fmt.Println("registering route: ", route)
		http.Handle(route, handler)
	}
	fmt.Println("Mote client listening on:", requestProcessor.listenPort)
	http.ListenAndServe(fmt.Sprintf(":%v", requestProcessor.listenPort), nil)
}
