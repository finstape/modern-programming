package main

import (
	"encoding/base64"
	"os"
	"testing"
)

type mockFileIO struct {
	readFileFunc  func(filename string) ([]byte, error)
	writeFileFunc func(filename string, data []byte, perm os.FileMode) error
}

func (m *mockFileIO) ReadFile(filename string) ([]byte, error) {
	if m.readFileFunc != nil {
		return m.readFileFunc(filename)
	}
	return nil, nil
}

func (m *mockFileIO) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if m.writeFileFunc != nil {
		return m.writeFileFunc(filename, data, perm)
	}
	return nil
}

func TestEncodeFile(t *testing.T) {
	mock := &mockFileIO{
		readFileFunc: func(filename string) ([]byte, error) {
			return []byte("Hello, World!"), nil
		},
		writeFileFunc: func(filename string, data []byte, perm os.FileMode) error {
			expectedData := base64.StdEncoding.EncodeToString([]byte("Hello, World!"))
			if string(data) != expectedData {
				t.Errorf("Unexpected encoded data. Expected: %s, Got: %s", expectedData, string(data))
			}
			return nil
		},
	}

	fileIO = mock
	defer func() { fileIO = &realFileIO{} }()

	encodeFile("input.txt", "output.txt")
}

func TestDecodeFile(t *testing.T) {
	mock := &mockFileIO{
		readFileFunc: func(filename string) ([]byte, error) {
			return []byte(base64.StdEncoding.EncodeToString([]byte("Hello, World!"))), nil
		},
		writeFileFunc: func(filename string, data []byte, perm os.FileMode) error {
			expectedData := "Hello, World!"
			if string(data) != expectedData {
				t.Errorf("Unexpected decoded data. Expected: %s, Got: %s", expectedData, string(data))
			}
			return nil
		},
	}

	fileIO = mock
	defer func() { fileIO = &realFileIO{} }()

	decodeFile("input.txt", "output.txt")
}
