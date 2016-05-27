package logg

import (
	"net/http"
	"strings"
	"time"
)

type Log struct {
	Server string
	Uri    string
	Param  string
	Method string
	Ip     string
}

func New() *Log {
	return &Log{}
}

func (l *Log) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	start := time.Now()

	next(w, r)

	CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========", RemoteIP(r))
	CLog("[TRAC] @@ 方法 @@: # %s #", r.Method)
	CLog("[TRAC] @@ 地址 @@: # %s #", r.RequestURI)
	CLog("[TRAC] @@ 参数 @@: # %s #", r.Form.Encode())
	CLog("[TRAC] @@ 用时 @@: ( %s )", time.Since(start))
	println("")
}

func RemoteIP(r *http.Request) string {
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

