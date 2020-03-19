package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	l4g "github.com/ezoic/log4go"
)

// HandleRequestForGVLVersion2 to return JSON sample of GVL Version 2 list for testing
func HandleRequestForGVLVersion2(rw http.ResponseWriter, req *http.Request) {
	gvlVersion2 := generatePrettifiedOutPutEN()
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(&gvlVersion2)
	_, err := rw.Write(b.Bytes())
	if err != nil {
		l4g.Error(err)
		http.Error(rw, "There was an error returning the cached value.", http.StatusInternalServerError)
	}
}
