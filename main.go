package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ezoic/ezcache"
	gvlcachev2 "github.com/ezoic/gvlcache/gvlcacheV2"
	"github.com/go-chi/chi"
)

// Reference documentation for this task can be found here
// https://app.asana.com/0/1163157365763573/1164482186726660/f

// ServerConfig is a struct representing an the configurations for this server
type ServerConfig struct {
	BindIPAddr string // The ip address for server to bind to.
	BindIPPort int    // The port for the server ot listen to.
	L4gConfig  string // Path of the configuration file for l4g
	CertPath   string // Path where cert public key is
	KeyPath    string // Path where cert private key is
}

func main() {
	fmt.Println("Starting server")
	// 1. Start memcache
	ezcache.InitializeMemcachedForRegion()
	// 2. Load configuration file for server

	// 3. Set up router object
	r := chi.NewRouter()
	r.Get("/", HandleRoot)
	r.Get("/GVLV2", gvlcachev2.HandleRequestForGVLVersion2)
	r.Post("/GVLV2Cache/bustCache", gvlcachev2.HandleRequestForBustingCache)
	log.Fatal(http.ListenAndServe(":8054", r))
}

// HandleRoot is a handler function for the root server that is used for testing
func HandleRoot(rw http.ResponseWriter, req *http.Request) {
	// l4g.Info("Received request from server.")
	rw.Write([]byte("Received request and processed the response for the root path"))
}
