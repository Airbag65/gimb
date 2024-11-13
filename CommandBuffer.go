package main

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
        Command: []string {},
    }
}

func (c *CommandBuffer) ExecuteCommand(s tcell.Screen) error {
    Command := strings.Join(c.Command[:], "")
    switch Command{
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

    c.Command = []string {}
    return nil
}


func (c *CommandBuffer) Add(char rune) {
    if char != ':'{
        c.Command = append(c.Command, string(char))
    }
}

func (c *CommandBuffer) DelKey() {
    c.Command = c.Command[:len(c.Command)-1]
    // pop(&c.Command)
}

func pop(list *[]string) string {
    l := len(*list)
    val := (*list)[l-1]
    *list = (*list)[:l-1] 
    return val
}


func (c *CommandBuffer) ToString() string {
    CmdString := strings.Join(c.Command[:], "")
    return CmdString
}




