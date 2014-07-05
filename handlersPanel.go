package main

import (
	//  "fmt"
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
	"github.com/gorilla/mux"
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
// Dynamic handlers
//====================================//
// Handle panel
func panelHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	vars := mux.Vars(r)
	debug("vars:", vars)
	debug("vars[\"page\"]:", vars["page"])
	debug("redirecting")
	http.Redirect(w, r, "/", http.StatusFound)
}
