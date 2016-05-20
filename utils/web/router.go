package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
	"1898/utils/log"
)

// @version 1
type Router struct {
	mux.Router
}

var track bool

func SetTrac(b bool) {
	track = b
}

func NewRoute() *Router {
	return &Router{}
}

func (r *Router) Add(method, path string, h http.Handler) *mux.Route {
	handler := Logger(h)
	return r.NewRoute().PathPrefix(path).Handler(handler).Methods(method)
}

func (r *Router) Get(path string, h http.HandlerFunc) *mux.Route {
	return r.Add("GET", path, h)
}

func (r *Router) Post(path string, h http.HandlerFunc) *mux.Route {
	return r.Add("POST", path, h)
}

func (r *Router) Delete(path string, h http.HandlerFunc) *mux.Route {
	return r.Add("DELETE", path, h)
}

func (r *Router) Put(path string, h http.HandlerFunc) *mux.Route {
	return r.Add("PUT", path, h)
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		log.CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========", IP(r))
		log.CLog("[TRAC] @@ 方法 @@: # %s #", r.Method)
		log.CLog("[TRAC] @@ 地址 @@: # %s #", r.RequestURI)
		log.CLog("[TRAC] @@ 参数 @@: # %s #", r.Form.Encode())
		log.CLog("[TRAC] @@ 用时 @@: ( %s )", time.Since(start))

		println("")
	})
}


func IP(r *http.Request) string {
	ipstr := r.Header.Get("X-Forwarded-For")
	ips := strings.Split(ipstr, ",")
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(r.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}
