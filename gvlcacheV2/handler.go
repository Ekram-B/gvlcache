package gvlcachev2

import (
	"bytes"
	"encoding/json"
	"net/http"

	l4g "github.com/ezoic/log4go"
)

// The cache key is independant of domain, user, and its generality servers as a way to sync callers as
// if there is only one cache key, then all callers will use the one cache key in order to retreive the
// value from the cache. The value's presence or not will determine if another call is made to IAB's
// server
const (
	gvlCacheKey string = "gvl-version2-key"
)

// HandleRequestForGVLVersion2 is the handler meant to process the GET request made for the GVL Version 2 List
func HandleRequestForGVLVersion2(rw http.ResponseWriter, req *http.Request) {
	l4g.Info("Hello")
	gvl := GVLVersionTwoValue{}
	isGVLInCache := gvl.getGVLVersionTwoValueFromCache()
	l4g.Info(isGVLInCache)
	if isGVLInCache == false {
		err := gvl.getGVLVersionTwoValueFromIABSource()
		if err != nil {
			http.Error(rw, "There was an error returning the vendor list from IAB's server.", http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(gvl)
	_, err := rw.Write(b.Bytes())
	if err != nil {
		l4g.Error(err)
		http.Error(rw, "There was an error returning the cached value.", http.StatusInternalServerError)
	}

}

// HandleRequestForBustingCache is the handler meant bust the cache if required
func HandleRequestForBustingCache(rw http.ResponseWriter, req *http.Request) {
	// 1. Bust the cache and remove all content from it
}
