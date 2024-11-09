package main

import "github.com/gdamore/tcell/v2"

func Red(s tcell.Screen) {
    usedStyle = tcell.StyleDefault.
            Foreground(tcell.ColorRed).
            Background(tcell.ColorBlack)
}

func Green(s tcell.Screen) {
    usedStyle = tcell.StyleDefault.
            Foreground(tcell.ColorGreen).
            Background(tcell.ColorBlack)
}

func Blue(s tcell.Screen) {
    usedStyle = tcell.StyleDefault.
            Foreground(tcell.ColorBlue).
            Background(tcell.ColorBlack)
}

func Help(s tcell.Screen) {
    defStyle := tcell.StyleDefault.
        Background(tcell.ColorBlack).
        Foreground(tcell.ColorWhite)
        
    s.Clear()
    PlaceText(s, 3, 3, "Help Menu -- Press Q to quit menu", defStyle)
    PlaceText(s, 3, 5, "Commands:", defStyle)
    PlaceText(s, 3, 6, "red", defStyle)
    PlaceText(s, 20, 6, "-- Change foreground color to red", defStyle)
    PlaceText(s, 3, 7, "green", defStyle)
    PlaceText(s, 20, 7, "-- Change foreground color to green", defStyle)
    PlaceText(s, 3, 8, "blue", defStyle)
    PlaceText(s, 20, 8, "-- Change foreground color to blue", defStyle)
    PlaceText(s, 3, 9, "q", defStyle)
    PlaceText(s, 20, 9, "-- Quit application", defStyle)
    PlaceText(s, 3, 10, "quit", defStyle)
    PlaceText(s, 20, 10, "-- Quit application", defStyle)
    PlaceText(s, 3, 11, "h", defStyle)
    PlaceText(s, 20, 11, "-- Help menu", defStyle)
    PlaceText(s, 3, 12, "help", defStyle)
    PlaceText(s, 20, 12, "-- Help menu", defStyle)
    s.Show() 

    for{
        ev := s.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Rune() == 'q' || ev.Rune() == 'Q'{ 
                s.Clear()
                return 
            }
        }
    }
}

func PlaceText(screen tcell.Screen, x int, y int, text string, style tcell.Style) {
    for i, char := range text {
        screen.SetContent(x + i, y, char, nil, style)
    } 
}
