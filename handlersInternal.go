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
// Slash redirect
func slashHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  debug("redirecting:", "/"+vars["request"]+"/")
  http.Redirect(w, r, "/"+vars["request"]+"/", http.StatusFound)
}
