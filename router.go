package main

import (
	"1898/dll"
	"1898/utils/log"
	"1898/utils/web"
	"net/http"
)

func HttpRun(addr string) {
	web.SetTrac(true)
	r := web.NewRoute()

	// 路由
	// user
	r.Post("/user/regkey", dll.CheckRegKey)
	r.Post("/user/register", dll.Register)
	r.Post("/user/login", dll.Login)

	// key
	r.Post("/key/getkey", dll.GetKey)

	// event
	r.Post("/event/new", dll.NewEvent)
	r.Post("/event/reg", dll.RegEvent)

	log.CLog("[TRAC] Server start listen on # %s #\n", addr)

	err := http.ListenAndServe(":"+addr, r)

	if err != nil {
		panic(err)
	}

}
