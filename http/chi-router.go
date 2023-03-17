package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chaiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chaiRouter{}
}

func (*chaiRouter) GET(uri string, f func(resp http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chaiRouter) POST(uri string, f func(resp http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chaiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
