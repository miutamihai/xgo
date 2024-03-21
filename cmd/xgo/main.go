package main

import (
	"os"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	args := os.Args[1:]

	if len(args) != 1 {
		panic("xgo requires exactly one argument")
	}

	clipboard.Write(clipboard.FmtText, []byte(args[0]))
}
