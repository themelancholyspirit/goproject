package main

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	addr string
}

func (s *Server) listenAndServe() {
	http.ListenAndServe(s.addr, nil)
}

func main() {

	http.HandleFunc("/users/", usersHandler)
	http.HandleFunc("/", mainEntryHandler)

	s := Server{
		addr: ":8080",
	}

	s.listenAndServe()

}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// var url string = r.RequestURI

	json.NewEncoder(w).Encode(Users)

}

func mainEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var msg map[string]string = map[string]string{
		"message": "This is the '/' endpoint...",
	}

	json.NewEncoder(w).Encode(msg)

}
