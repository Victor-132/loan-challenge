package httpserver

import "net/http"

type HttpServer[t any] interface {
	On(method, path string, callback func(t) error)
	Listen(port int)
	Test(req *http.Request) (*http.Response, error)
}
