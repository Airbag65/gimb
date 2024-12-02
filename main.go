package main

import (
	// "fmt"
	// "log"
 //    "os"
	"github.com/gdamore/tcell/v2"
    // "strings"
)
var usedStyle tcell.Style = tcell.StyleDefault.
        Background(tcell.ColorBlack).
        Foreground(tcell.ColorWhite)



func main(){
    window := NewWindow()
    window.Draw()
 //    screen, err := tcell.NewScreen()
	//
 //    file := NewFile(string(os.Args[1]))
 //    file.FileContent = append(file.FileContent, "Test from the program")
 //    if err != nil{
 //        log.Fatalf("%+v", err)
 //    }
 //    if err := screen.Init(); err != nil {
 //        log.Fatalf("%+v", err) 
 //    }
	//
 //    buffer := NewBuffer()
 //    commandMode := false
 //    insertMode := false
	//
 //    quit := func() {
	// 	maybePanic := recover()
	// 	screen.Fini()
	// 	if maybePanic != nil {
	// 		panic(maybePanic)
 //            // return
	// 	}
	// }
	// defer quit()
	//
 //   
 //    var mode string
 //    for {
 //        if commandMode{
 //            mode = "COMMAND"
 //        }else if (insertMode){
 //            mode = "INSERT"
 //        }else{
 //            mode = "NORMAL"
 //        }
 //        screen.Clear()
	//
 //        ev := screen.PollEvent()
	//
 //        switch ev := ev.(type) {
 //        case *tcell.EventKey:
 //            if ev.Key() == tcell.KeyEscape{
 //                continue
 //            }
 //            if commandMode && ev.Key() == tcell.KeyEnter{
 //                mode = "NORMAL"
 //                commandMode = false
 //                err := buffer.ExecuteCommand(screen, file)
 //                if err != nil {
 //                    return
 //                }
 //            }
 //            if commandMode && ev.Key() == tcell.KeyBackspace2{
 //                buffer.DelKey()
 //                // draw(screen, file, buffer, mode)
 //                continue
	//
 //            }
 //            if ev.Rune() == ':'{
 //                mode = "COMMAND"
 //                commandMode = true
 //            }
 //            if commandMode{
 //                buffer.Add(ev.Rune())
 //                PlaceText(screen, 5, 50, fmt.Sprintf("Command=> :%s", buffer.ToString()), usedStyle)
 //            }    
 //        }
        // width, _ := screen.Size()
        // screen.SetStyle(usedStyle)
        // for i := 0; i < width; i++ {
        //     screen.SetContent(width - i, 49, '_', nil, usedStyle)
        //     screen.SetContent(width - i, 5, '_', nil, usedStyle)
        // }
        // for i, line := range file.FileContent {
        //     lineNum := fmt.Sprintf("%d|", i+1)
        //     lineNum = strings.Repeat(string(' '), 4 - len(lineNum)) + lineNum
        //     PlaceText(screen, 2, 6 + i, lineNum, tcell.StyleDefault.Foreground(tcell.ColorGray))
        //     PlaceText(screen, 7, 6 + i, line, tcell.StyleDefault.Foreground(tcell.ColorWhite))
        // }
        // PlaceText(screen, (width / 2) - 2, 2, "GIMB", usedStyle)
        // PlaceText(screen, (width / 2) - ((len(mode) + 6) / 2), 4, fmt.Sprintf("-- %s --", mode), usedStyle)
        // PlaceText(screen, (width / 2) - ((len(mode) + 2) / 2), 3, fmt.Sprintf("\"%s\"", file.FilePath), usedStyle)
        // PlaceText(screen, 2, 3, "Press ESC, or use command q/quit, to exit", usedStyle)
        // PlaceText(screen, 5, 51, "Use command h/help for list of commands", usedStyle)
        // screen.Show()
    // }
}
