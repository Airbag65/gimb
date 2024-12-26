package utils

import "github.com/gdamore/tcell/v2"

type Cursor struct {
	Char           rune
	CoordX, CoordY int
    BlinkShowing   bool
    timer          int16
    Style          tcell.Style
}

func CreateCursor(x, y int, char rune) *Cursor {
	return &Cursor{
		Char:         char,
		CoordX:       x,
		CoordY:       y,
        BlinkShowing: false,
        timer:        0,
        Style:        tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite),
	}
}

// TODO: On hold for now
// func (c *Cursor) Blink() {
//     if c.BlinkShowing{
//         c.Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
//     } else {
//         c.Style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
//     }
//
//     c.timer++
//     if c.timer == 1 {
//         c.BlinkShowing = !c.BlinkShowing
//         c.timer = 0
//     }
// }
