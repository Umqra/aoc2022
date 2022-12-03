package day3

import (
	"github.com/Umqra/aoc2022/internal"
	"testing"
)

func TestRucksack(t *testing.T) {
	r := parseRucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	elements := r.findCommonElements()
	total := 0
	for _, element := range elements {
		total += element.Cost()
	}
	if total != 38 {
		t.Fatalf("%v != 38", total)
	}
}

func TestThreeRucksacks(t *testing.T) {
	r1 := parseRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	r2 := parseRucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	r3 := parseRucksack("PmmdzqPrVvPwwTWBwg")
	elements := findCommonElements([]Rucksack{r1, r2, r3})
	total := 0
	for _, element := range elements {
		total += element.Cost()
	}
	if total != 18 {
		t.Fatalf("%v != 18", total)
	}
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v", internal.Eval("day3.a.in", Solve1))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v", internal.Eval("day3.a.in", Solve2))
}
