package main

import (
	"fmt"
	//  "log"
	// File handling
	//  "os"
	//  "io/ioutil"
	// Configuration files handling
	//  "encoding/json"
	// HTTP Server, FastCGI
	//  "net"
	"net/http"
	//  "net/http/fcgi"
	// Router
	//  "github.com/gorilla/mux"
	// Path decoding, markdown/BBCode
	//  "regexp"
	// Webpage generation
	//  "html/template"
)

//####################################//
// Types & Structs
//====================================//
// Types
// FIXME

//====================================//
// Structs
// FIXME

//####################################//
// JSON API handlers
//====================================//
// Handle JSON API
// FIXME
func exampleHandlerJSON(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Header().Set("Content-Type", "application/json")
	debug("Sending JSON example")
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}
