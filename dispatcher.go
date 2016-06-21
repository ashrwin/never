package main

import "os"

var event = make(chan string)

func emptyEventChan() {
	for {
		select {
		case _ = <-event:

		default:
			return
		}
	}
}
func startDispatcher() {
	watch(options.Root)
	var process *os.Process
	if build() {
		process = run()
	}

	for _ = range event {
		if build() {

			emptyEventChan()
			if process == nil {
				process = run()
			} else {
				process = restart(process)
			}

		}
	}

}
