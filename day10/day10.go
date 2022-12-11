package day10

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Command interface {
	Execute(c *Cpu, ticks chan<- Cpu)
}

type Cpu struct {
	register int
}

type Computer struct {
	cpu          Cpu
	ticks        <-chan Cpu
	instructions chan<- Command
}

type Noop struct{}
type Add struct{ value int }

func (n *Noop) Execute(c *Cpu, ticks chan<- Cpu) {
	ticks <- *c
}

func (a *Add) Execute(c *Cpu, ticks chan<- Cpu) {
	ticks <- *c
	ticks <- *c
	c.register += a.value
}

func NewComputer() Computer {
	instructions := make(chan Command)
	ticks := make(chan Cpu)
	computer := Computer{
		cpu:          Cpu{register: 1},
		ticks:        ticks,
		instructions: instructions,
	}
	go func() {
		for i := range instructions {
			i.Execute(&computer.cpu, ticks)
		}
		close(ticks)
	}()
	return computer
}

func EnterCommands(f *os.File, computer Computer) {
	reader := internal.NewFileReader(f)
	for reader.Scan() {
		lineReader := internal.NewStringReader(reader.ParseString()).SetDelimiter(" ")
		var command Command
		commandString := lineReader.ScanString()
		if commandString == "noop" {
			command = &Noop{}
		} else if commandString == "addx" {
			command = &Add{value: lineReader.ScanInt()}
		} else {
			panic(fmt.Sprintf("unknown command string: %v", commandString))
		}
		computer.instructions <- command
	}
	close(computer.instructions)
}

func Solve1(f *os.File) interface{} {
	computer := NewComputer()
	go EnterCommands(f, computer)

	strength := 0
	cycle := 0
	for cpu := range computer.ticks {
		cycle++
		if cycle%40 == 20 {
			strength += cpu.register * cycle
		}
	}
	return strength
}

func Solve2(f *os.File) interface{} {
	computer := NewComputer()
	go EnterCommands(f, computer)

	for i := 0; i < 6; i++ {
		var line [40]byte
		for s := range line {
			line[s] = ' '
			if cpu := <-computer.ticks; cpu.register-1 <= s && s <= cpu.register+1 {
				line[s] = '#'
			}
		}
		fmt.Printf("%v\n", string(line[:]))
	}
	return nil
}
