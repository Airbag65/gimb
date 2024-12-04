package main

import (
	"fmt"
	"gimb_out/utils"
	"log"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen        tcell.Screen
	Style         tcell.Style
	CommandBuffer utils.CommandBuffer
	File          utils.File
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
	width, _ := w.Screen.Size()
	for i := 0; i < width; i++ {
		w.Screen.SetContent(width-i, 49, '_', nil, tcell.StyleDefault)
		w.Screen.SetContent(width-i, 5, '_', nil, tcell.StyleDefault)
		w.Screen.SetContent(width-i, 50, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGray))
	}
	_ = len(w.Mode.String()) + 2
	// Lägga texten där den ska vara!
	utils.PlaceText(w.Screen, width-15, 50, w.Mode.String(), tcell.StyleDefault.Background(tcell.ColorGray))
}

func (w *Window) DrawCommandLine() {
	if w.Mode == utils.CommandMode {
		utils.PlaceText(w.Screen, 5, 50, fmt.Sprintf("Command => :%s", w.CommandBuffer.ToString()), tcell.StyleDefault.Background(tcell.ColorGray))
	}
}

func (w *Window) HandleEvent() error {
	event := w.Screen.PollEvent()
	switch ev := event.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			w.Mode = utils.NormalMode
			return fmt.Errorf("Exit")
		} else if ev.Rune() == ':' {
			w.Mode = utils.CommandMode
		} else if ev.Rune() == 'i' {
			w.Mode = utils.InsertMode
		}
	}

	return nil
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
		w.DrawFrame()
		w.DrawCommandLine()
		if err := w.HandleEvent(); err != nil {
			return
		}
		w.Screen.Show()
	}
}
