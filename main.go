package main

import (
	"os"
)

func main() {
    defer func(){
        recover()
    }()
	path := string(os.Args[1])
	if path == "" {
		path = "tmp.txt"
    }
	window := NewWindow(path)
	window.Draw()

}
