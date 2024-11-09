package main

import (
    "fmt"
    // "log"
	"github.com/gdamore/tcell/v2"
)

type CommandBuffer struct {
    Command string
}


func NewBuffer() *CommandBuffer {
    return &CommandBuffer{
        Command: "",
    }
}

func (c *CommandBuffer) ExecuteCommand(s tcell.Screen) error {
    switch c.Command{
    case "red":
        Red(s)
    case "green":
        Green(s)
    case "blue":
        Blue(s)
    case "q":
        return fmt.Errorf("Quit")
    case "quit":
        return fmt.Errorf("Quit")
    case "h":
        Help(s)
    case "help":
        Help(s)
    case "def":
        Default(s)
    case "yellow":
        Yellow(s)
    default:
        PlaceText(s, 5, 50, "Unrecognized Command", tcell.StyleDefault.Foreground(tcell.ColorRed))
    }

    c.Command = ""
    return nil
}


func (c *CommandBuffer) Add(char rune) {
    if char != ':'{
        c.Command += string(char)
    }
}

func (c *CommandBuffer) DelKey() {
    // TODO:  Fixa den här skiten
    if len(c.Command) != 0{
        c.Command = c.Command[:len(c.Command) - 1]
    }
}
