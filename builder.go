package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() bool {
	cmd := exec.Command("go", "build", "-o", options.BuildFolder, options.Root)

	buildLog("\n\nBuilding...\n\n")

	stderr, err := cmd.StderrPipe()
	checkError(err)

	stdout, err := cmd.StdoutPipe()
	checkError(err)

	err = cmd.Start()
	checkError(err)

	io.Copy(os.Stdout, stdout)
	stderrbuff, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()

	if err != nil {
		buildLog(string(stderrbuff))
		return false
	}
	buildLog("Build Successfull...\n\n")

	return true
}
