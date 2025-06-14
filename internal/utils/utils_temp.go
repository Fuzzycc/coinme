package utils

import (
	"os"
)

// --- --- --- ---
// Read
// --- --- --- ---

// A wrapper for IORead. Syntatic sugar.
func IOReadTemp(f *os.File) (string, error) {
	data, err := IORead(f)
	if err != nil {
		return "", err
	}

	return data, nil
}

// --- --- --- ---
// Write
// --- --- --- ---

// An orthogonal to IOWrite. Takes an *os.File rather than io.Reader.
func IOWriteTemp(f *os.File, input string) (n int, e error) {
	n, e = f.WriteString(input)
	if e != nil {
		return 0, e
	}

	return n, nil
}

// --- --- --- ---
// Create
// --- --- --- ---

func CreateTemp(dirname string, name string) (*os.File, error) {
	dname, err := os.MkdirTemp("", dirname)
	if err != nil {
		return nil, err
	}

	f, err := os.CreateTemp(dname, name)
	if err != nil {
		return nil, err
	}

	return f, nil
}
