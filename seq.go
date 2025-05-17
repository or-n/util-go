package util

import (
	"encoding/gob"
	"io"
	"os"
)

func ToSeq(w io.Writer, data any) error {
	return gob.NewEncoder(w).Encode(data)
}

func FromSeq(w io.Reader, data any) error {
	return gob.NewDecoder(w).Decode(data)
}

func Save(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return ToSeq(file, data)
}

func Load(filename string, data any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return FromSeq(file, data)
}
