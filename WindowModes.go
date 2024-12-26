package main

import (
	"gimb_out/utils"
    "github.com/gdamore/tcell/v2"
)

func (w *Window) HandleCommandMode(event tcell.Event) error{
    switch ev := event.(type) {
    case *tcell.EventKey:
        if ev.Key() == tcell.KeyEnter {
            err := w.CommandBuffer.ExecuteCommand(w.Screen, &w.File)
            if err != nil {
                return err
            }
            w.Mode = utils.NormalMode
            w.DrawFrame()
        } else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
            w.CommandBuffer.DelKey()
            w.DrawFrame()
            w.DrawCommandLine()
        } else if ev.Key() == tcell.KeyEscape {
            w.Mode = utils.NormalMode
            w.CommandBuffer.Clear()
            w.DrawFrame()
        } else {
            w.CommandBuffer.Add(ev.Rune())
            w.DrawFrame()
            w.DrawCommandLine()
        }
    }
    return nil
}

func (w *Window) HandleInsertMode(event tcell.Event) error {
    switch ev := event.(type) {
    case *tcell.EventKey:
        if ev.Key() == tcell.KeyEscape {
            w.Mode = utils.NormalMode
            w.DrawFrame()
        }
    }
    return nil
}
