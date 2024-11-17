package main

import (
    // "fmt"
    "bufio"
    "os"
    "strings"
)


type File struct {
    FilePath string 
    FileContent []string
}


func NewFile(Path string) *File {
    _, err := os.Stat(Path)
    if os.IsNotExist(err) {
        file, err := os.Create(Path)
        if err != nil{
            return nil
        }
        defer file.Close()
       
        return &File{
            FilePath: Path,
            FileContent: []string{},
        }       
    }
    file, err := os.Open(Path)
    if err != nil{
        return nil
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    content := []string{}
    for scanner.Scan() {
        content = append(content, scanner.Text())
    }
    
    return &File{
        FilePath: Path,
        FileContent: content,
    }
}

func (file *File) SaveFile() {
    compressedFileContent := strings.Join(file.FileContent[:], "\n")
    // panic(compressedFileContent)
    f, err := os.OpenFile(file.FilePath, os.O_WRONLY, 0644)
    if err != nil{
        return
    }
    defer f.Close() 

    // writer := bufio.NewWriter(f)
    f.Write([]byte(compressedFileContent))
    // writer.Flush()

}
