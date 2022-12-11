package day09

import (
	"github.com/Umqra/aoc2022/internal"
	"testing"
)

func TestExample1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day9.test.in", Solve1))
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day9.a.in", Solve1))
}

func TestExample2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day9.test.in", Solve2))
}

func TestExampleLarge2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day9.test2.in", Solve2))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day9.a.in", Solve2))
}
