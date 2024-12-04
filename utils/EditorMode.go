package utils

type EditorMode int8

const (
	NormalMode EditorMode = iota
	InsertMode
	CommandMode
)

func (em EditorMode) String() string {
	switch em {
	case NormalMode:
		return "NORMAL"
	case InsertMode:
		return "INSERT"
	case CommandMode:
		return "COMMAND"
	}
	return "unknown"
}
