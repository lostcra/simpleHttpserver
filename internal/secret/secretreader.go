package secret

import (
	"errors"
	"os"
	"strings"
)

type Reader interface {
	ReadFile() (string, error)
}

type FileReader struct {
	Path string
}

func New(path string) (*FileReader, error) {
	if path == "" {
		return nil, errors.New("path variable is empty")
	}

	return &FileReader{Path: path}, nil
}

func (sfr FileReader) ReadFile() (string, error) {
	password, err := os.ReadFile(sfr.Path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(password)), nil
}
