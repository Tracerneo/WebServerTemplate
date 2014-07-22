package main

import (
	"log"
	// File handling
	"io/ioutil"
	"path/filepath"
	// Configuration files handling
	"encoding/json"
)

//####################################//
// Configuration handling
//====================================//
// Server Configuration
//  Read Only, loaded at start of server
type ServerConf struct {
	Config SConf `json:"config"`
}
type SConf struct {
	Version float64     `json:"version"`
	Server  SConfServer `json:"server"`
	Path    SConfPath   `json:"path"`
	Admin   SConfAdmin  `json:"admin"`
	DB      SConfDB     `json:"database"`
}
type SConfServer struct {
	FCGI   bool   `json:"fcgi"`
	Type   string `json:"type"`
	Addr   string `json:"address"`
	Port   string `json:"port"`
	Socket string `json:"socket"`
}
type SConfPath struct {
	Domain   SConfDomain `json:"domain"`
	BasePath string      `json:"basepath"`
	Panel    string      `json:"panel"`
}
type SConfDomain struct {
	Main   string `json:"main"`
	Static string `json:"static"`
	Images string `json:"images"`
}
type SConfAdmin struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}
type SConfDB struct {
	Type string `json:"type"`
	Addr string `json:"address"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

//====================================//
// Parse Configuration
func parseConfig() (*ServerConf) {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("Error: ParseConfig():", err)
	}
	c := new(ServerConf)
	if err := json.Unmarshal(configFile, c); err != nil {
		log.Println("Warning: ParseConfig():", err)
		return c
	}
	return c
}

//====================================//
// Configure various things
func configure(c *ServerConf) {
	c.Config.Path.BasePath = filepath.Clean(c.Config.Path.BasePath)
	if c.Config.Path.BasePath == "/" || c.Config.Path.BasePath == "." {
		c.Config.Path.BasePath = ""
	}

	// FIXME add more things

}
