package main

import (
	"1898/dll"
	"1898/utils/log"
	"net/http"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"1898/middleware/logg"
)

func HttpRun(addr string) {

	r := mux.NewRouter()

	r.HandleFunc("/login", dll.Login).Methods("POST", "GET")
	r.HandleFunc("/register", dll.Register).Methods("POST")
	r.HandleFunc("/regkey", dll.CheckRegKey).Methods("POST")

	r.HandleFunc("/user/info", dll.GetUserByID)

	// key
	r.HandleFunc("/key/getkey", dll.GetKey).Methods("POST", "GET")

	// event
	r.HandleFunc("/event/new", dll.NewEvent).Methods("POST")
	r.HandleFunc("/event/edit", dll.EditEvent).Methods("POST")
	r.HandleFunc("/event/info", dll.EventInfo).Methods("POST")
	r.HandleFunc("/event/reg", dll.RegEvent).Methods("POST")
	r.HandleFunc("/event/list", dll.EventList).Methods("POST")
	r.HandleFunc("/event/del", dll.DelEvent).Methods("POST")
	r.HandleFunc("/event/cancel", dll.CancelEvent).Methods("POST")

	n := negroni.New()
	n.Use(logg.New())
	n.UseHandler(r)

	log.CLog("[TRAC] Server start listen on # %s #\n", addr)

	err := http.ListenAndServe(":"+addr, n)

	if err != nil {
		panic(err)
	}

}
