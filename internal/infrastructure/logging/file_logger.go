package logging

import "os"

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{file: file}, nil
}

func (f *FileLogger) Log(message string) error {
	_, err := f.file.WriteString(message + "\n")
	return err
}

func (f *FileLogger) Close() error {
	return f.file.Close()
}