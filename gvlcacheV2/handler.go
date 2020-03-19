package gvlcachev2

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
	// 1. Create an instance of the GVL object to store content into
	gvl := GVLVersionTwoValue{}
	// 2. NOW WE CAN CHECK THE CACHE!! If the control flow has reached this point, we can assume the request args exists and was used properly to build a valid cache key.
	// TODO write logic for passing in the key and returning a value
	isGVLInCache := gvl.getGVLVersionTwoValueFromCache()

	if isGVLInCache == false {
		// A) If we haven't found the entry in the cache, then the gvl.Value member will still be empty for the type, and we have to set it
		// 1. Validate value using the constraints specified in the technical documentation -> We can assume the IAB return is correct, but this is done to make sure our processing logic succeeds
		// 2. If validated to have met the constaints, then we return the marshalled result back to the caller
		err := gvl.getGVLVersionTwoValueFromIABSource()
		if err != nil {
			http.Error(rw, "There was an error returning the vendor list from IAB's server.", http.StatusInternalServerError)
			return
		}
	}
	// B) If we have found an entry in the cache, then the gvl.Value member won't be empty, and we can return it

	rw.Header().Add("Content-Type", "application/json")
	// Set the status code to OK since this is the desired result
	rw.WriteHeader(http.StatusOK)
	// Encode the part of the GVL struct that we care about
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(gvl)
	_, err := rw.Write(b.Bytes())
	if err != nil {
		log.Println(err)
		http.Error(rw, "There was an error returning the cached value.", http.StatusInternalServerError)
	}

}

// HandleRequestForBustingCache is the handler meant bust the cache if required
func HandleRequestForBustingCache(rw http.ResponseWriter, req *http.Request) {
	// 1. Bust the cache and remove all content from it
}
