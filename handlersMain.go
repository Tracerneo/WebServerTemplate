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
// Handle something
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	vars := mux.Vars(r)
	debug("vars:", vars)
	debug("vars[\"subject\"]:", vars["subject"])
	if(vars["subject"] != "") {
		fmt.Fprintln(w, "Hello", vars["subject"]+"!")
	}	else {
		fmt.Fprintln(w, "Hello World!")
	}
}

//####################################//
// Static handlers
//====================================//
// Handle static content
func staticHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	debug("redirecting")
	http.Redirect(w, r, "/", http.StatusFound)
}
