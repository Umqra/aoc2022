package day15

import (
	"github.com/Umqra/aoc2022/internal"
	"os"
	"testing"
)

func TestExample1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day15.test.in", func(file *os.File) interface{} {
		return Solve1(file, 10)
	}))
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day15.a.in", func(file *os.File) interface{} {
		return Solve1(file, 2000000)
	}))
}

func TestExample2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day15.test.in", func(file *os.File) interface{} {
		return Solve2(file, 20)
	}))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day15.a.in", func(file *os.File) interface{} {
		return Solve2(file, 4000000)
	}))
}
