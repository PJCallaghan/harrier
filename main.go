package main

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

func main() {
	mux := initMux()
	fmt.Println("Server started on port 80. Visit http://localhost:80")
	http.ListenAndServe(":80", mux)
}
