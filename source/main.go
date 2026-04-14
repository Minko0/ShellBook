package main

import "source/ui"

func main() {
	tui := ui.New(GetCommands())
	tui.Run()
}
