package day11

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int64
	operation func(x int64) int64
	divider   int64
	throw     func(x int64) int
	inspected int
}

func parseLine(line string, delimiter string) []string {
	tokens := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), delimiter)
	result := make([]string, len(tokens))
	for i := 0; i < len(tokens); i++ {
		result[i] = strings.Trim(tokens[i], " ")
	}
	return result
}

func parseOperation(operation []string) func(x int64) int64 {
	var self bool
	var argument int
	if operation[4] == "old" {
		self = true
	} else {
		var err error
		argument, err = strconv.Atoi(operation[4])
		if err != nil {
			panic(fmt.Sprintf("unable to parse integer from argument (%v): %v", operation, err))
		}
	}
	if operation[3] == "+" {
		return func(x int64) int64 {
			if self {
				return x + x
			}
			return x + int64(argument)
		}
	} else if operation[3] == "*" {
		return func(x int64) int64 {
			if self {
				return x * x
			}
			return x * int64(argument)
		}
	} else {
		panic(fmt.Sprintf("unknown operation: %v", operation[3]))
	}
}

func parseThrowAction(test []string, trueAction []string, falseAction []string) (int64, func(x int64) int) {
	divider, err := strconv.Atoi(test[2])
	if err != nil {
		panic(fmt.Sprintf("unable to parse integer from test argument: %v", err))
	}
	trueMonkey, err := strconv.Atoi(trueAction[3])
	if err != nil {
		panic(fmt.Sprintf("unable to parse integer from true argument: %v", err))
	}
	falseMonkey, err := strconv.Atoi(falseAction[3])
	if err != nil {
		panic(fmt.Sprintf("unable to parse integer from false argument: %v", err))
	}
	return int64(divider), func(x int64) int {
		if x%int64(divider) == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}

type Monkeys []Monkey

func (m Monkeys) Play(divide bool, pacer int64) {
	for i := 0; i < len(m); i++ {
		items := m[i].items
		m[i].inspected += len(items)
		m[i].items = make([]int64, 0)
		for _, item := range items {
			worry := m[i].operation(item)
			if divide {
				worry /= 3
			} else {
				worry %= pacer
			}
			next := m[i].throw(worry)
			m[next].items = append(m[next].items, worry)
		}
	}
}

func parseMonkeys(f *os.File) Monkeys {
	reader := internal.NewFileReader(f).SetDelimiter("\n\n")
	monkeys := make(Monkeys, 0)
	for reader.Scan() {
		monkeyReader := internal.NewStringReader(reader.ParseString())
		_ = monkeyReader.ScanString()
		items := make([]int64, 0)
		for _, token := range parseLine(monkeyReader.ScanString(), ",") {
			item, err := strconv.Atoi(token)
			if err != nil {
				panic(fmt.Sprintf("unable to parse integer from token %v: %v", token, err))
			}
			items = append(items, int64(item))
		}
		operation := parseOperation(parseLine(monkeyReader.ScanString(), " "))
		test := parseLine(monkeyReader.ScanString(), " ")
		trueAction := parseLine(monkeyReader.ScanString(), " ")
		falseAction := parseLine(monkeyReader.ScanString(), " ")
		divider, throwAction := parseThrowAction(test, trueAction, falseAction)
		monkeys = append(monkeys, Monkey{
			items:     items,
			operation: operation,
			throw:     throwAction,
			divider:   divider,
		})
	}
	return monkeys
}

func Solve1(f *os.File) interface{} {
	monkeys := parseMonkeys(f)
	for i := 0; i < 20; i++ {
		monkeys.Play(true, 0)
	}
	inspected := make([]int, len(monkeys))
	for i := 0; i < len(monkeys); i++ {
		inspected[i] = monkeys[i].inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}

func Solve2(f *os.File) interface{} {
	monkeys := parseMonkeys(f)
	pacer := int64(1)
	for _, monkey := range monkeys {
		pacer *= monkey.divider
	}
	for i := 0; i < 10000; i++ {
		monkeys.Play(false, pacer)
	}
	inspected := make([]int, len(monkeys))
	for i := 0; i < len(monkeys); i++ {
		inspected[i] = monkeys[i].inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return int64(inspected[0]) * int64(inspected[1])
}
