package day01

import (
	"github.com/Umqra/aoc2022/internal"
	"testing"
)

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day1.a.in", Solve1))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day1.a.in", Solve2))
}
