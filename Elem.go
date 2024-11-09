package main

import "github.com/gdamore/tcell/v2"


type Element struct{
    coordX, coordY int
    char rune
}

func NewElement(char rune, x, y int) *Element{
    return &Element{
        char: char,
        coordX: x,
        coordY: y,
    }
}

func (e *Element) PlaceElement (s tcell.Screen){
    s.SetContent(e.coordX, e.coordY, e.char, nil, tcell.StyleDefault)
}
