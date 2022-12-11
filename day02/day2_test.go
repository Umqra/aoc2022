package day02

import (
	"github.com/Umqra/aoc2022/internal"
	"testing"
)

func TestScore_Update(t *testing.T) {
	score := Score(0)
	score.Update(Rock, Paper)
	score.Update(Paper, Rock)
	score.Update(Scissors, Scissors)
	if score != 15 {
		t.Fatalf("score must be equal to 15")
	}
}

func TestElement_ElementForResult(t *testing.T) {
	if Rock.ElementForResult(Win).GameResult(Rock) != Win {
		t.Fatalf("invalid game result")
	}
	if Rock.ElementForResult(Draw).GameResult(Rock) != Draw {
		t.Fatalf("invalid game result")
	}
	if Rock.ElementForResult(Lose).GameResult(Rock) != Lose {
		t.Fatalf("invalid game result")
	}
}

func TestScore_UpdateByOutcome(t *testing.T) {
	score := Score(0)
	score.UpdateByOutcome(Rock, Draw)
	if score != 4 {
		t.Fatalf("%v != 4", score)
	}
	score.UpdateByOutcome(Paper, Lose)
	score.UpdateByOutcome(Scissors, Win)
	if score != 12 {
		t.Fatalf("%v != 12", score)
	}
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day2.a.in", Solve1))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day2.a.in", Solve2))
}
