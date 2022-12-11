package day5

import (
	"github.com/Umqra/aoc2022/internal"
	"os"
	"strings"
	"unicode"
)

type Table [][]byte

func (t Table) MoveSlice(count int, from, to int) {
	slice := t[from][len(t[from])-count:]
	t[from] = t[from][:len(t[from])-count]
	t[to] = append(t[to], slice...)
}

func (t Table) MoveOneByOne(count int, from, to int) {
	for ; count > 0; count-- {
		t.MoveSlice(1, from, to)
	}
}

func (t Table) Top() string {
	result := make([]byte, len(t))
	for i := 0; i < len(t); i++ {
		if len(t[i]) == 0 {
			result[i] = ' '
		} else {
			result[i] = t[i][len(t[i])-1]
		}
	}
	return string(result)
}

func parseTable(table string) Table {
	lines := strings.Split(table, "\n")

	last := lines[len(lines)-1]
	position := make([]int, len(last))
	id := 0
	for s := 0; s < len(last); s++ {
		position[s] = id
		if s > 0 && last[s] == ' ' && unicode.IsDigit(rune(last[s-1])) {
			id++
		}
	}
	stacks := make([][]byte, id+1)
	for i := len(lines) - 2; i >= 0; i-- {
		for s := 0; s < len(lines[i]); s++ {
			if unicode.IsLetter(rune(lines[i][s])) {
				stacks[position[s]] = append(stacks[position[s]], lines[i][s])
			}
		}
	}
	return stacks
}

func Solve1(f *os.File) interface{} {
	reader := internal.NewFileReader(f).SetDelimiter("\n\n")
	table := parseTable(reader.ScanString())
	operations := strings.Trim(reader.ScanString(), "\n")
	for _, operation := range strings.Split(operations, "\n") {
		parser := internal.NewStringReader(operation).SetDelimiter(" ")
		parser.ScanToken("move")
		count := parser.ScanInt()
		parser.ScanToken("from")
		from := parser.ScanInt() - 1
		parser.ScanToken("to")
		to := parser.ScanInt() - 1
		table.MoveOneByOne(count, from, to)
	}

	return table.Top()
}

func Solve2(f *os.File) interface{} {
	reader := internal.NewFileReader(f).SetDelimiter("\n\n")
	table := parseTable(reader.ScanString())
	operations := strings.Trim(reader.ScanString(), "\n")
	for _, operation := range strings.Split(operations, "\n") {
		parser := internal.NewStringReader(operation).SetDelimiter(" ")
		parser.ScanToken("move")
		count := parser.ScanInt()
		parser.ScanToken("from")
		from := parser.ScanInt() - 1
		parser.ScanToken("to")
		to := parser.ScanInt() - 1
		table.MoveSlice(count, from, to)
	}

	return table.Top()
}
