package main

//####################################//
// Packages importing
//====================================//
// Global packages
import (
	"path/filepath"
	"runtime"
	"strconv"
)

//####################################//
// Code
//====================================//
// Debug function
func debug(v ...interface{}) {
	if flagDebug {
		pc, filename, line, _ := runtime.Caller(1)
		filename = filepath.Base(filename)
		funcptr := runtime.FuncForPC(pc)
		funcname := funcptr.Name()
		logline := []interface{}{"Debug:", filename + ":" + strconv.Itoa(line), funcname + ":"}
		v = append(logline, v...)
		Stddebug.Println(v...)
	}
}
