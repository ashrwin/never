package main

import (
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

// color package is not safe to be used across multipe
// concurrent goroutines
var mu sync.Mutex
var buildLog = func(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	color.Yellow(format, a...)

}
var watchLog = func(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	color.Magenta(format, a...)
}

var runLog = func(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	color.White(format, a...)
}
var mainLog = func(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	color.White(format, a...)
}

var errorLog = func(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	color.Cyan(format, a...)
}

var fatal = func(err error) {
	mu.Lock()
	defer mu.Unlock()
	log.Println(err.Error())
	os.Exit(1)
}
