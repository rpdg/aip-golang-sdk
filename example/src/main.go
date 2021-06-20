package main

import (
	"aip-face-sdk/example/src/core"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var router = mux.NewRouter()

var port = ":8086"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	core.RenderHome(w, r, "index")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	core.RenderLogin(w, r, "login")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	core.RenderLogout(w, r, "logout")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	core.RenderAdmin(w, r, "admin")
}

func init() {
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/admin", accountHandler)

	router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css/"))))
	router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js/"))))
	router.PathPrefix("/static/images/").Handler(http.StripPrefix("/static/images/", http.FileServer(http.Dir("static/images/"))))
}

func main() {

	//启动服务
	log.Printf("Server is running at http://localhost%s/. Press Ctrl+C to stop.", port)

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    3600 * time.Second,
		WriteTimeout:   3600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
