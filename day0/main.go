package main

import (
	"log"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// this part is additional, if we want the home page to be displayed if and only if the request url path exactly matches `/`
	// http.NotFound replies to the request with an HTTP 404 not found error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, write := w.Write([]byte("Hello World from home page"))
	if write != nil {
		return
	}
}

func main() {
	// create a router
	mux := http.NewServeMux()

	// register the home handler based on a pattern
	mux.HandleFunc("/", HandleHome)

	// create a server that uses mux
	log.Println("Starting server on port 8000...")
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
