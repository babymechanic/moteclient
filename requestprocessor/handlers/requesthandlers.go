package handlers

import (
	"net/http"
)

type RequestHandlers map[string]http.Handler

var requestHandlers = RequestHandlers{}

func RegisterHandler(route string, handler http.Handler) {
	requestHandlers[route] = handler
}

func Handlers() RequestHandlers {
	return requestHandlers
}
