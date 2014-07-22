package main

//####################################//
// Packages importing
//====================================//
// Global packages
import (
	"flag"
	"fmt"
	"log"
	// File handling
	"os"
	// HTTP Server, FastCGI
	"net"
	"net/http"
	"net/http/fcgi"
	// Router
	"github.com/gorilla/mux"
)

//====================================//
// Local packages
import ()

//####################################//
// Global variables
//====================================//
// Engine variables
var c *ServerConf // Parsed config structure
var Stddebug *log.Logger

// Flags
// FIXME add flags
var flagDebug bool
var flagQuiet bool

//####################################//
// Code
//====================================//
// Init code
func init() {
	// FIXME init flags
	const (
		flagDef_Debug = false
		flagDescDebug = "   -debug   enable debug"
		flagDef_Quiet = false
		flagDescQuiet = "-q -quiet   suppress output"
	)
	flag.BoolVar(&flagDebug, "debug", false, flagDescDebug)
	flag.BoolVar(&flagQuiet, "q", false, flagDescQuiet)
	flag.BoolVar(&flagQuiet, "quiet", false, flagDescQuiet)

	//----------------------------------//
	// Help
	// FIXME update help
	flag.Usage = func() {
		usage := `Usage: %s [options]
Options:
  ` + flag.Lookup("q").Usage + `
  ` + flag.Lookup("debug").Usage + `
  No more options for now.

MIT, BSD or something. There is no help.
`
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	//----------------------------------//
}

// Main code
func main() {
	flag.Parse()
	engineInit()
	c, _ := parseConfig()
	debug(c)
	configure(c)

	router := mux.NewRouter()
	debug("c.Config.Path.BasePath =", c.Config.Path.BasePath)

	// FIXME add routes
	// Serve static content
	router.
		//Host(c.Config.Path.Domain.Static).
		PathPrefix("/static/").
		Handler(http.FileServer(http.Dir("htdocs/")))
	router.
		//Host(c.Config.Path.Domain.Images).
		PathPrefix("/src/").
		Handler(http.FileServer(http.Dir("htdocs/img/")))
	router.
		//Host(c.Config.Path.Domain.Images).
		PathPrefix("/thumb/").
		Handler(http.FileServer(http.Dir("htdocs/img/")))

	// Main subrouter
	r := router.
		//FIXME Host("{subdomain:[a-z]+}."+c.Config.Path.Domain.Main).
		//FIXME Host(c.Config.Path.Domain.Main).
		PathPrefix(c.Config.Path.BasePath + "/").
		Subrouter()

	//----------------------------------//
	// Slash redirect
	r.HandleFunc("/{request}", slashHandler)

	//----------------------------------//
	// Handle dynamic pages
	// Handle panel
	panel := r.PathPrefix("/" + c.Config.Path.Panel).Subrouter()
	panel.HandleFunc("/", panelHandler)
	panel.HandleFunc("/{page}", panelHandler)

	// Handle custom webpages
	// FIXME
	r.HandleFunc("/", helloWorldHandler)
	r.HandleFunc("/hello/", helloWorldHandler)
	r.HandleFunc("/hello/{subject}", helloWorldHandler)

	// Handle JSON API
	// FIXME
	j := r.PathPrefix("/json").Subrouter()
	j.HandleFunc("/example.json", exampleHandlerJSON)

	//----------------------------------//
	// Main HTTP handler
	http.Handle("/", router)

	//----------------------------------//
	// Start webserver
	if !c.Config.Server.FCGI && c.Config.Server.Type == "unix" {
		err := "\nI'm not sure using unix socket with http server is good idea." +
			"\nConsider changing type to 'tcp' or enabling 'fcgi' in config."
		log.Fatalln("Error: main():", err)
	}
	var listener net.Listener
	var err error
	switch c.Config.Server.Type {
	case "unix":
		listener, err = net.Listen(c.Config.Server.Type, c.Config.Server.Socket)
		if err != nil {
			log.Fatalln("Error: main():", err)
		}
	case "tcp", "tcp4", "tcp6":
		listener, err = net.Listen(c.Config.Server.Type,
			net.JoinHostPort(c.Config.Server.Addr, c.Config.Server.Port))
		if err != nil {
			log.Fatalln("Error: main():", err)
		}
	default:
		log.Fatalln("Error: main(): unsupported type", c.Config.Server.Type)
	}
	if c.Config.Server.FCGI {
		log.Println("Info: Starting FastCGI")
		if err := fcgi.Serve(listener, nil); err != nil {
			log.Fatalln("Error: main():", err)
		}
	} else {
		log.Println("Info: Starting HTTP Server")
		if err := http.Serve(listener, nil); err != nil {
			log.Fatalln("Error: main():", err)
		}
	}
}
