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

	// user
	r.HandleFunc("/login", dll.Login).Methods("POST")
	r.HandleFunc("/rootlogin", dll.RootLogin).Methods("POST")
	r.HandleFunc("/register", dll.Register).Methods("POST")
	r.HandleFunc("/regkey", dll.CheckRegKey).Methods("POST")

	r.HandleFunc("/user/newroot", dll.CreateRoot)
	r.HandleFunc("/user/editpwd", dll.EditPassword)
	r.HandleFunc("/user/edit", dll.EditUser)
	r.HandleFunc("/user/info", dll.GetUserByID)

	// friends
	r.HandleFunc("/friend/add", dll.AddFriend)
	r.HandleFunc("/friend/list", dll.FriendsList)
	r.HandleFunc("/friend/del", dll.DelFriend)

	//message
	r.HandleFunc("/msg/push", dll.PushMsg)
	r.HandleFunc("/msg/pull", dll.PullMsg)
	r.HandleFunc("/msg/read", dll.ReadMsg)

	// key
	r.HandleFunc("/key/getkey", dll.GetKey).Methods("POST")

	// event
	r.HandleFunc("/event/new", dll.NewEvent).Methods("POST")
	r.HandleFunc("/event/edit", dll.EditEvent).Methods("POST")
	r.HandleFunc("/event/info", dll.EventInfo).Methods("POST")
	r.HandleFunc("/event/reg", dll.RegEvent).Methods("POST")
	r.HandleFunc("/event/list", dll.EventList).Methods("POST")
	r.HandleFunc("/event/del", dll.DelEvent).Methods("POST")
	r.HandleFunc("/event/cancel", dll.CancelEvent).Methods("POST")

	// news
	r.HandleFunc("/news/add", dll.AddNews)
	r.HandleFunc("/news/edit", dll.EditNews)
	r.HandleFunc("/news/list", dll.NewsList)
	r.HandleFunc("/news/del", dll.DelNews)

	n := negroni.New()
	n.Use(logg.New())
	n.UseHandler(r)

	log.CLog("[TRAC] Server start listen on # %s #\n", addr)

	err := http.ListenAndServe(":"+addr, n)

	if err != nil {
		panic(err)
	}

}
