package internal

import (
	"fmt"
	"os"
)

func Eval(filename string, eval func(file *os.File) interface{}) interface{} {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("unable to open file: %w", err))
	}
	defer f.Close()
	return eval(f)
}
