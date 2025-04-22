package service

import (
	"fmt"
	"net/http"
)

func initMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	return mux
}

func NewServer() StartableServer {
	return &httpServer{}
}

func start() error {
	mux := initMux()
	fmt.Println("Server started on port 80. Visit http://localhost:80")
	err := http.ListenAndServe(":80", mux)
	return err
}

type StartableServer interface {
	Start() error
}

type httpServer struct{}

func (s *httpServer) Start() error {
	return start()
}
