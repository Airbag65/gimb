package utils

import (
	"bufio"
	"os"
	"strings"
)

type FileManager struct {
	FilePath    string
	FileContent [][]rune
    Cursor      Cursor
}

func NewFile(Path string) *FileManager {
	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		file, err := os.Create(Path)
		if err != nil {
			return nil
		}
		defer file.Close()

		return &FileManager{
			FilePath:    Path,
			FileContent: [][]rune{},
            Cursor:      *CreateCursor(0, 0, ' '),
		}
	}
	file, err := os.Open(Path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	content := [][]rune{}
	for scanner.Scan() {
        str := scanner.Text()
        content = append(content, []rune(str))
	}

	return &FileManager{
		FilePath:    Path,
		FileContent: content,
        Cursor:      *CreateCursor(0, 0, content[0][0]),
	}
}

func (file *FileManager) SaveFile() {
    strSlice := []string{}
    for _, line := range file.FileContent{
        strSlice = append(strSlice, string(line))
    }
	compressedFileContent := strings.Join(strSlice[:], "\n")
	f, err := os.OpenFile(file.FilePath, os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	f.Write([]byte(compressedFileContent))

}
