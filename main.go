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
    var path string
    if len(os.Args) > 0 {
        path = "tmp.txt"
    } else {
        path = string(os.Args[1])
    }
	window := NewWindow(path)
	window.Draw()
}
