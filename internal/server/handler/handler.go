package handler

import "net/http"

type HttpHandler interface {
	GetUrlPattern() string
	GetHandler() func(http.ResponseWriter, *http.Request)
}
