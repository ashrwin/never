package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

type Options struct {
	BuildFolder string
	Root        string
	Delay       time.Duration
	ArgsForProg []string
	UserOptions
}

//options which can be provided by user
type UserOptions struct {
	IgnoredFolders    []string `json:"ignoredFolders"`
	IgnoredExtensions []string `json:"ignoredExtensions"`
}

func LoadOptions() Options {
	var configFile = "./never.conf.json"
	buildFolder := "./tmp/runner"
	if runtime.GOOS == "windows" {
		buildFolder += ".exe"
	}
	var uoptions UserOptions
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		uoptions = UserOptions{
			[]string{"./tmp"},
			make([]string, 0),
		}
	} else {
		raw, _ := ioutil.ReadFile(configFile)
		err := json.Unmarshal(raw, &uoptions)
		checkError(err)
		uoptions.IgnoredFolders = append(uoptions.IgnoredFolders, "./tmp")
	}
	args := make([]string, 0)
	//args[0] is this program's name
	for _, k := range os.Args[1:] {
		args = append(args, k)
	}
	return Options{
		buildFolder,
		".",
		500,
		args,
		uoptions,
	}

}
