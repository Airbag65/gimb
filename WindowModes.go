package main

import (
	"gimb_out/utils"
	"slices"

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
            return nil
        } else if ev.Key() == tcell.KeyEnter{
            w.File.FileContent = slices.Insert(w.File.FileContent, w.File.Cursor.CoordY + 1, []rune{})
            w.File.FileContent[w.File.Cursor.CoordY + 1] = append(w.File.FileContent[w.File.Cursor.CoordY + 1], w.File.FileContent[w.File.Cursor.CoordY][w.File.Cursor.CoordX:]...)
            w.File.FileContent[w.File.Cursor.CoordY] = w.File.FileContent[w.File.Cursor.CoordY][:w.File.Cursor.CoordX]
            w.File.Cursor.CoordY++
            w.File.Cursor.CoordX = 0
        } else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
            w.File.FileContent[w.File.Cursor.CoordY] = append(w.File.FileContent[w.File.Cursor.CoordY][:w.File.Cursor.CoordX - 1], w.File.FileContent[w.File.Cursor.CoordY][w.File.Cursor.CoordX:]...)
            w.File.Cursor.CoordX--
            // if len(w.File.FileContent[w.File.Cursor.CoordY]) >= 1{
            //     w.File.FileContent = append(w.File.FileContent[:w.File.Cursor.CoordY - 1], w.File.FileContent[w.File.Cursor.CoordY])
            // } else {
            // }
        } else {
            w.File.FileContent[w.File.Cursor.CoordY] = slices.Insert(w.File.FileContent[w.File.Cursor.CoordY], w.File.Cursor.CoordX, ev.Rune())
            w.File.Cursor.CoordX++
        }
        w.DrawFileContent()
    }
    return nil
}


func (w *Window) HandleNormalMode(event tcell.Event) error {
    switch ev := event.(type){
    case *tcell.EventKey:
        if ev.Rune() == 'h'{
            if w.File.Cursor.CoordX != 0{
                w.File.Cursor.CoordX--
            }
        } else if ev.Rune() == 'l'{
            if w.File.Cursor.CoordX < len(w.File.FileContent[w.File.Cursor.CoordY]){
                w.File.Cursor.CoordX++
            }
        } else if ev.Rune() == 'j' {
            if w.File.Cursor.CoordY < len(w.File.FileContent) - 1{
                w.File.Cursor.CoordY++
                if w.File.Cursor.CoordX > len(w.File.FileContent[w.File.Cursor.CoordY]){
                   w.File.Cursor.CoordX = len(w.File.FileContent[w.File.Cursor.CoordY]) 
                }
            }
        } else if ev.Rune() == 'k' {
            if w.File.Cursor.CoordY != 0 {
                w.File.Cursor.CoordY--
                if w.File.Cursor.CoordX > len(w.File.FileContent[w.File.Cursor.CoordY]){
                   w.File.Cursor.CoordX = len(w.File.FileContent[w.File.Cursor.CoordY]) 
                }
            }
        }
    }

    if w.File.Cursor.CoordX < len(w.File.FileContent[w.File.Cursor.CoordY]){
        w.File.Cursor.Char = w.File.FileContent[w.File.Cursor.CoordY][w.File.Cursor.CoordX]
    } else {
        w.File.Cursor.Char = ' '
    }
    return nil
}
