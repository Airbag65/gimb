package utils

import (
	"github.com/gdamore/tcell/v2"
)

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
        // display
    } else {
        c.InBlink = true
        // display
    }
}
