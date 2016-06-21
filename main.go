package main

var options Options

func main() {
	mainLog("Starting never...")
	options = LoadOptions()
	startDispatcher()
}
