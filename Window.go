package main

import (
	"fmt"
	"gimb_out/utils"
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen        tcell.Screen
	Style         tcell.Style
	CommandBuffer utils.CommandBuffer
	File          utils.FileManager
	Mode          utils.EditorMode
}

func NewWindow(filepath string) *Window {
	s, err := tcell.NewScreen()
	if err != nil {
		panic("Now that's real bad....")
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	return &Window{
		Screen:        s,
		Style:         tcell.StyleDefault,
		CommandBuffer: *utils.NewBuffer(),
		File:          *utils.NewFile(filepath),
		Mode:          utils.NormalMode,
	}
}

func (w *Window) DrawFrame() {
	width, height := w.Screen.Size()
	for i := 0; i < width; i++ {
		w.Screen.SetContent(width-i, height - 4, '_', nil, tcell.StyleDefault)
		w.Screen.SetContent(width-i, 5, '_', nil, tcell.StyleDefault)
		w.Screen.SetContent(width-i, height - 3, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGray))
	}
	l := len(w.Mode.String())
	FillerRight := strings.Repeat("-", 12-l)
	ToPlaceMode := fmt.Sprintf("--%s%s", w.Mode.String(), FillerRight)
	ToPlaceFilename := fmt.Sprintf("--%s--", w.File.FilePath)

	utils.PlaceText(w.Screen, width - 12, height - 3,ToPlaceMode, tcell.StyleDefault.Background(tcell.ColorGray))
	utils.PlaceText(w.Screen, 1, 4, ToPlaceFilename, w.Style)
}

func (w *Window) DrawCommandLine() {
	_, height := w.Screen.Size()
	if w.Mode == utils.CommandMode {
		utils.PlaceText(w.Screen, 5, height - 3, fmt.Sprintf("Command => :%s", w.CommandBuffer.ToString()), tcell.StyleDefault.Background(tcell.ColorGray))
	}
}

func (w *Window) HandleEvent(e chan error){
	event := w.Screen.PollEvent()
	switch ev := event.(type) {
	case *tcell.EventKey:
		if ev.Rune() == ':' && w.Mode == utils.NormalMode{
			w.Mode = utils.CommandMode
            w.DrawCommandLine()
		} else if ev.Rune() == 'i' && w.Mode == utils.NormalMode {
			w.Mode = utils.InsertMode
            w.DrawFrame()
            break
		}
        if w.Mode == utils.CommandMode {
            if err := w.HandleCommandMode(event); err != nil {
                e <- err
            }
        } else if w.Mode == utils.InsertMode {
            if err := w.HandleInsertMode(event); err != nil {
                e <- err
            }
        } else if w.Mode == utils.NormalMode {
            if err := w.HandleNormalMode(event); err != nil {
                e <- err
            }
        }
    }
    e <- nil
}

func (w *Window) DrawFileContent(){
    for i, line := range w.File.FileContent {
        lineNum := fmt.Sprintf("%d|", i+1)
        lineNum = strings.Repeat(string(' '), 4 - len(lineNum)) + lineNum
        utils.PlaceText(w.Screen, 2, 6 + i, lineNum, tcell.StyleDefault.Foreground(tcell.ColorGray))
        utils.PlaceText(w.Screen, 7, 6 + i, string(line), tcell.StyleDefault.Foreground(tcell.ColorWhite))
    }
}

func (w *Window) DrawCursor(){
    utils.PlaceText(w.Screen, w.File.Cursor.CoordX + 7, w.File.Cursor.CoordY + 6, string(w.File.Cursor.Char), w.File.Cursor.Style)
}

func (w *Window) Draw() {
	quit := func() {
		maybePanic := recover()
		w.Screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	for {
        w.Screen.Clear()
		w.DrawFrame()
		w.DrawCommandLine()
        w.DrawFileContent()
        w.DrawCursor()

		w.Screen.Show()

        c := make(chan error)
        go w.HandleEvent(c)
        err := <- c
        if err != nil {
            panic(err)
        }
	}
}
