package day07

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type ExplorerDir struct {
	name string
}

type ExplorerFile struct {
	name string
	size int
}

type Explorer struct {
	stats map[string][]interface{}
	path  string
}

func (e *Explorer) RegisterCd(command string) {
	e.path = path.Clean(path.Join(e.path, command))
}

func (e *Explorer) RegisterLs(items []interface{}) {
	e.stats[e.path] = items
}

func (e *Explorer) analyze(d string, result map[string]int) int {
	if size, ok := result[d]; ok {
		return size
	}
	size := 0
	for _, element := range e.stats[d] {
		if dir, ok := element.(ExplorerDir); ok {
			size += e.analyze(path.Join(d, dir.name), result)
		} else if file, ok := element.(ExplorerFile); ok {
			size += file.size
		} else {
			log.Panicf("unexpected element type: %v", element)
		}
	}
	result[d] = size
	return size
}

func (e *Explorer) Analyze() map[string]int {
	sizes := make(map[string]int)
	for d := range e.stats {
		sizes[d] = e.analyze(d, sizes)
	}
	return sizes
}

type ShellReader internal.Reader

func NewShellReader(r internal.Reader) ShellReader {
	r.Scan()
	return ShellReader(r)
}

func (r ShellReader) parseShellCommand() (bool, string, []string) {
	command := internal.Reader(r).ParseString()[2:]
	if strings.HasPrefix(command, "cd") {
		return r.Scan(), command, nil
	} else if strings.HasPrefix(command, "ls") {
		output := make([]string, 0)
		for r.Scan() {
			line := internal.Reader(r).ParseString()
			if line[0] == '$' {
				return true, command, output
			}
			output = append(output, line)
		}
		return false, command, output
	}
	panic(fmt.Sprintf("unexpected command: %v", command))
}

func runShell(f *os.File) Explorer {
	explorer := Explorer{path: "/", stats: map[string][]interface{}{}}
	reader := NewShellReader(internal.NewFileReader(f))
	for {
		hasNext, command, output := reader.parseShellCommand()
		if command[:2] == "cd" {
			explorer.RegisterCd(command[3:])
		} else if command[:2] == "ls" {
			elements := make([]interface{}, 0)
			for _, element := range output {
				tokens := strings.Split(element, " ")
				if tokens[0] == "dir" {
					elements = append(elements, ExplorerDir{name: tokens[1]})
				} else {
					size, err := strconv.Atoi(tokens[0])
					if err != nil {
						log.Panicf("unexpected ls output: %v", element)
					}
					elements = append(elements, ExplorerFile{size: size, name: tokens[1]})
				}
			}
			explorer.RegisterLs(elements)
		}
		if !hasNext {
			break
		}
	}
	return explorer
}

func Solve1(f *os.File) interface{} {
	explorer := runShell(f)
	stat := explorer.Analyze()
	total := 0
	for _, size := range stat {
		if size <= 100000 {
			total += size
		}
	}
	return total
}

func Solve2(f *os.File) interface{} {
	explorer := runShell(f)
	stat := explorer.Analyze()
	total := stat["/"]
	best := total
	for _, size := range stat {
		if total-size+30000000 <= 70000000 && size < best {
			best = size
		}
	}
	return best
}
