package router

import "net/http"

type Router interface {
	GET(uri string, f func(resp http.ResponseWriter, request *http.Request))
	POST(uri string, f func(resp http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
