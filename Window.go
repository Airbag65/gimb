package main

import (
	"fmt"
	"gimb_out/utils"
	"log"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
    Screen tcell.Screen
    Style tcell.Style
    CommandBuffer utils.CommandBuffer
}

func NewWindow() *Window{
    s, err := tcell.NewScreen()
    if err != nil {
       panic("Now that's real bad....") 
    }     
    if err := s.Init(); err != nil {
        log.Fatalf("%+v", err) 
    }
    return &Window{
        Screen: s,
        Style: tcell.StyleDefault,
        CommandBuffer: *utils.NewBuffer(),
    }
}

func (w *Window) DrawFrame(){
    width, _ := w.Screen.Size()
    for i := 0; i < width; i++{
        w.Screen.SetContent(width - i, 49, '_', nil, tcell.StyleDefault)
        w.Screen.SetContent(width - i, 5, '_', nil, tcell.StyleDefault)
    }
}

func (w *Window) DrawCommandLine(){
    utils.PlaceText(w.Screen, 5, 50, fmt.Sprintf("Command => :%s", w.CommandBuffer.ToString()), w.Style)
}



func (w *Window) Draw(){
    w.DrawFrame()
    w.DrawCommandLine()
    w.Screen.Show()
}
