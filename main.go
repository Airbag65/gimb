package main

import (
	"fmt"
	"os"
)

func main() {
    defer func(){
        e := recover()
        fmt.Print(e)
    }()
	path := string(os.Args[1])
	if path == "" {
		path = "tmp.txt"
    }
	window := NewWindow(path)
	window.Draw()
    // fmt.Print(window.File.FileContent)
}
