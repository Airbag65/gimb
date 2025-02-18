package utils

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"strings"
)

type CommandBuffer struct {
	Command []string
}

func NewBuffer() *CommandBuffer {
	return &CommandBuffer{
		Command: []string{},
	}
}

func (c *CommandBuffer) ExecuteCommand(s tcell.Screen, f *FileManager) error {
	Command := strings.Join(c.Command[:], "")
	switch Command {
	case "q":
		return fmt.Errorf("Quit")
	case "quit":
		return fmt.Errorf("Quit")
	case "h":
		Help(s)
	case "help":
		Help(s)
	case "w":
		f.SaveFile()
	case "wq":
		f.SaveFile()
		return fmt.Errorf("Save and quit")
	default:
		PlaceText(s, 5, 50, "Unrecognized Command", tcell.StyleDefault.Foreground(tcell.ColorRed))
	}

    c.Clear()
	return nil
}

func (c *CommandBuffer) Add(char rune) {
	if char != ':' {
		c.Command = append(c.Command, string(char))
	}
}

func (c *CommandBuffer) DelKey() {
	if len(c.Command) > 0 {
		c.Command = c.Command[:len(c.Command)-1]
	}
}

func (c *CommandBuffer) Clear() {
    c.Command = []string{}
}

func (c *CommandBuffer) ToString() string {
	CmdString := strings.Join(c.Command[:], "")
	return CmdString
}
