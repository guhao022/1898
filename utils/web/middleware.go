package web

import "net/http"

type middleware struct {
	mux      http.Handler // for the router/mux. (Gorilla, httprouter, etc)
	handlers []http.HandlerFunc // middleware funcs to run.
}

func M() *middleware {
	return &middleware{handlers: make([]http.HandlerFunc, 0, 0)}
}

// Add adds a variable number of handlers using variadic arguments.
func (m *middleware) Add(h ...http.HandlerFunc) {
	m.handlers = append(m.handlers, h...)
}

// AddMux adds our mux to run.
func (m *middleware) AddMux(mux http.Handler) {
	m.mux = mux
}

// So we can satisfy the http.Handler interface.
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, h := range m.handlers {
		h.ServeHTTP(w, r)
	}
	m.mux.ServeHTTP(w, r)
}