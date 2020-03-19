package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// 1. Load configuration file for server
	// 2. Set up router object
	r := chi.NewRouter()
	r.Get("/", HandleRoot)
	r.Get("/v2/vendor-list.json", HandleRequestForGVLVersion2)
	log.Fatal(http.ListenAndServe(":8055", r))
}

// HandleRoot is a handler function for the root server that is used for testing
func HandleRoot(rw http.ResponseWriter, req *http.Request) {
	// l4g.Info("Received request from server.")
	rw.Write([]byte("Received request and processed the response for the root path"))
}
