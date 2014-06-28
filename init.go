//  vim: set ts=2 sw=2 tw=0 et :

package main

import (
  // Logging
  "log"
  // File handling
  "os"
  "io"
  "path/filepath"
)

func engineInit() {
  execdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatalln("Error: Init():", err)
  }
  if err = os.Chdir(execdir); err != nil {
    log.Fatalln("Error: Init():", err)
  }
  initLogger()
}

func initLogger() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  //log.SetFlags(log.Ldate | log.Ltime)
  logFile, err := os.OpenFile("log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
  if err != nil {
    log.Fatalln("Error: InitLogger():", err)
  } else {
    log.SetOutput(io.MultiWriter(os.Stderr, logFile))
    log.Println("Info: Logger initialized")
  }
  if flagDebug {
    Stddebug = log.New(os.Stderr, "", log.Ldate | log.Ltime)
  }
}
