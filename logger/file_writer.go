package logger

import (
	"os"
)

type FileWriter struct {
	file *os.File
}

func NewFileWriter(filename string) *FileWriter {
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return &FileWriter{file: f}
}

func (f *FileWriter) Write(message string) {
	f.file.WriteString(message + "\n")
}
