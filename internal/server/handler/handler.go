package handler

import "net/http"

type HttpHandler interface {
	GetUrlPattern() string
	GetHandler() func(http.ResponseWriter, *http.Request)
}

//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
//w.Header().Set("Access-Control-Allow-Credentials", "true")
//w.Header().Set("Access-Control-Allow-Origin", "*")
//w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
