package utils

import (
    "io/ioutil"
    "os"
)

// ReadFile reads the contents of a file.
func ReadFile(filePath string) ([]byte, error) {
    return ioutil.ReadFile(filePath)
}

// WriteFile writes data to a file.
func WriteFile(filePath string, data []byte) error {
    return ioutil.WriteFile(filePath, data, 0644)
}

// CreateDir creates a directory if it doesn't already exist.
func CreateDir(path string) error {
    return os.MkdirAll(path, 0755)
}
