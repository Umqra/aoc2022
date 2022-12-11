package day10

import (
	"github.com/Umqra/aoc2022/internal"
	"testing"
)

func TestExample1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day10.test.in", Solve1))
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day10.a.in", Solve1))
}

func TestExample2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day10.test.in", Solve2))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day10.a.in", Solve2))
}
