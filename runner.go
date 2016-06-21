package main

import (
	"io"
	"os"
	"os/exec"
)

func run() *os.Process {

	cmd := exec.Command(options.BuildFolder, options.ArgsForProg...)

	stderr, err := cmd.StderrPipe()
	checkError(err)
	stdout, err := cmd.StdoutPipe()
	checkError(err)

	err = cmd.Start()
	checkError(err)

	runLog("Running process")
	printSeperatorLine(runLog)

	go io.Copy(os.Stdout, stderr)
	go io.Copy(os.Stdout, stdout)

	return cmd.Process

}

func restart(process *os.Process) *os.Process {

	runLog("Killing process PID : %d \n\n", process.Pid)
	process.Kill()
	return run()
}
