package main

import (
	"github.com/gdamore/tcell/v2"
)

var style tcell.Style = tcell.StyleDefault.
        Foreground(tcell.ColorBlack).
        Background(tcell.ColorWhite)

type Cursor struct{
   Char rune
   InBlink bool
   CoordX, CoordY int
}

func CreateCursor(x, y int, char rune) *Cursor{
    return &Cursor{
        Char: char,
        InBlink: true,
        CoordX: x,
        CoordY: y,
    }
}


func (c Cursor) Blink(s tcell.Screen) {
    if c.InBlink{
        c.InBlink = false
        style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
    } else {
        c.InBlink = true
        style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
    }
}
