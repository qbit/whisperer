package main

import "github.com/qbit/whisperer/cmd"

func main() {
	root := cmd.Root()
	root.Execute()
}
