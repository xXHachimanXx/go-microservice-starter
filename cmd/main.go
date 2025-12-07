package main

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world!!!"))
}

func main() {
	s := &Server{addr: ":8080"}

	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal("Deu erro aqui ze")
	}
}
