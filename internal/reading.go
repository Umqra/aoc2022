package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Reader struct{ *bufio.Scanner }

func MustOpenFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return f
}

func NewStringReader(input string) Reader {
	return Reader{bufio.NewScanner(strings.NewReader(input))}
}

func NewFileReader(file *os.File) Reader {
	return Reader{bufio.NewScanner(file)}
}

func (reader Reader) SetDelimiter(delimiter string) Reader {
	b := []byte(delimiter)
	reader.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) >= len(b) {
			for i := range data[0 : len(data)-len(b)+1] {
				if bytes.Compare(data[i:i+len(b)], b) == 0 {
					return i + len(b), data[0:i], nil
				}
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	return reader
}

func (reader Reader) ScanString() string {
	if !reader.Scan() {
		panic(fmt.Errorf("unexpected eof"))
	}
	return reader.ParseString()
}

func (reader Reader) ScanInt() int {
	if !reader.Scan() {
		panic(fmt.Errorf("unexpected eof"))
	}
	return reader.ParseInt()
}

func (reader Reader) ScanToken(pattern string) string {
	if !reader.Scan() {
		panic(fmt.Errorf("unexpected eof"))
	}
	return reader.ParseToken(pattern)
}

func (reader Reader) ParseString() string {
	return reader.Text()
}

func (reader Reader) ParseInt() int {
	i, err := strconv.Atoi(reader.Text())
	if err != nil {
		panic(fmt.Errorf("unable to parse int: %w", err))
	}
	return i
}

func (reader Reader) ParseToken(pattern string) string {
	s := reader.Text()
	match, err := regexp.MatchString("^"+pattern+"$", s)
	if err != nil {
		panic(fmt.Errorf("regexp pattern (%v) matching error: %w", pattern, err))
	}
	if !match {
		panic(fmt.Errorf("token pattern mismatch: %v NOT LIKE %v", s, pattern))
	}
	return s
}
