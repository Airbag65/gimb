package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)
var usedStyle tcell.Style = tcell.StyleDefault.
        Background(tcell.ColorBlack).
        Foreground(tcell.ColorWhite)

func main(){
    screen, err := tcell.NewScreen()
    if err != nil{
        log.Fatalf("%+v", err)
    }
    if err := screen.Init(); err != nil {
        log.Fatalf("%+v", err) 
    }

    buffer := NewBuffer()
    commandMode := false

    // screen.SetStyle(tcell.StyleDefault)
    // screen.Clear()
    
    quit := func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
            return
		}
	}
	defer quit()

   
    for {
        screen.Clear()

        ev := screen.PollEvent()

        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape{
                return
            }
            if commandMode && ev.Key() == tcell.KeyEnter{
                commandMode = false
                err := buffer.ExecuteCommand(screen)
                if err != nil {
                    return
                }
            }
            if commandMode && ev.Key() == tcell.KeyBackspace2{
                buffer.DelKey()
            }
            if ev.Rune() == ':'{
                commandMode = true
            }
            if commandMode{
                buffer.Add(ev.Rune())
                PlaceText(screen, 5, 50, fmt.Sprintf("Command=> :%s", buffer.Command), usedStyle)
            }    
        }
        width, _ := screen.Size()
        screen.SetStyle(usedStyle)
        // PlaceText(screen, 10, 10, fmt.Sprintf("%d * %d", a, b), usedStyle)
        for i := 0; i < width; i++{
            screen.SetContent(width - i, 49, '_', nil, usedStyle)
            screen.SetContent(width - i, 4, '_', nil, usedStyle)
        }
        PlaceText(screen, 2, 2, "Welcome to the thing!", usedStyle)
        PlaceText(screen, 2, 3, "Press ESC, or use command q/quit, to exit", usedStyle)
        PlaceText(screen, 5, 51, "Use command h/help for list of commands", usedStyle)
        screen.Show()

    }
}
