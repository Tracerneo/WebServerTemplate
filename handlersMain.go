package main

import (
	//  "fmt"
	"log"
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
	"html/template"
)

//####################################//
// Types & Structs
//====================================//
// Types
// FIXME

//====================================//
// Structs
// FIXME
type Hello struct {
	Subject string
}

//####################################//
// Dynamic handlers
//====================================//
// Handle something
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	vars := mux.Vars(r)
	debug("vars:", vars)
	debug("vars[\"subject\"]:", vars["subject"])
	h := Hello{}
	if vars["subject"] != "" {
		h.Subject = vars["subject"]
	} else {
		h.Subject = "World"
	}
	t, err := template.ParseFiles("template/hello.tpl")
	if err != nil {
		log.Fatalln("Error: main():", err)
	}
	t.Execute(w, h)
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
