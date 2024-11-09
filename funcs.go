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

func Default(s tcell.Screen) {
    usedStyle = tcell.StyleDefault.
            Foreground(tcell.ColorWhite).
            Background(tcell.ColorBlack)
}

func Yellow(s tcell.Screen) {
    usedStyle = tcell.StyleDefault.
            Foreground(tcell.ColorYellow).
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

    PlaceText(s, 3, 8, "green", defStyle)
    PlaceText(s, 20, 8, "-- Change foreground color to green", defStyle)

    PlaceText(s, 3, 10, "blue", defStyle)
    PlaceText(s, 20, 10, "-- Change foreground color to blue", defStyle)

    PlaceText(s, 3, 12, "def", defStyle)
    PlaceText(s, 20, 12, "-- Change foreground color to defult color", defStyle)

    PlaceText(s, 3, 14, "yellow", defStyle)
    PlaceText(s, 20, 14, "-- Change foreground color to yellow", defStyle)

    PlaceText(s, 3, 16, "q", defStyle)
    PlaceText(s, 20, 16, "-- Quit application", defStyle)

    PlaceText(s, 3, 18, "quit", defStyle)
    PlaceText(s, 20, 18, "-- Quit application", defStyle)

    PlaceText(s, 3, 20, "h", defStyle)
    PlaceText(s, 20, 20, "-- Help menu", defStyle)

    PlaceText(s, 3, 22, "help", defStyle)
    PlaceText(s, 20, 22, "-- Help menu", defStyle)

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
