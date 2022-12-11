package day02

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Element int

const (
	Rock     Element = 0
	Paper    Element = 1
	Scissors Element = 2
)

type Outcome int

const (
	Win  Outcome = 6
	Draw Outcome = 3
	Lose Outcome = 0
)

func (a Element) Cost() int {
	return int(a) + 1
}

func (a Element) GameResult(b Element) Outcome {
	if a == b {
		return Draw
	} else if a == (b+1)%3 {
		return Win
	}
	return Lose
}

func (a Element) ElementForResult(outcome Outcome) Element {
	if outcome == Draw {
		return a
	} else if outcome == Lose {
		return (a + 2) % 3
	} else {
		return (a + 1) % 3
	}
}

func MustParseElement(s string) (element Element) {
	switch s {
	case "A", "X":
		element = Rock
	case "B", "Y":
		element = Paper
	case "C", "Z":
		element = Scissors
	default:
		panic(fmt.Errorf("unable to parse element string: %v", s))
	}
	return
}

func MustParseOutcome(s string) (outcome Outcome) {
	switch s {
	case "X":
		outcome = Lose
	case "Y":
		outcome = Draw
	case "Z":
		outcome = Win
	default:
		panic(fmt.Errorf("unable to parse outcome string: %v", s))
	}
	return
}

type Score int

func (score *Score) Update(opponent, player Element) {
	*score = Score(int(*score) + player.Cost() + int(player.GameResult(opponent)))
}

func (score *Score) UpdateByOutcome(opponent Element, outcome Outcome) {
	*score = Score(int(*score) + opponent.ElementForResult(outcome).Cost() + int(outcome))
}

func Solve1(f *os.File) interface{} {
	score := Score(0)
	reader := internal.NewFileReader(f)
	for reader.Scan() {
		line := internal.NewStringReader(reader.ParseString()).SetDelimiter(" ")
		opponent := MustParseElement(line.ScanToken("A|B|C"))
		player := MustParseElement(line.ScanToken("X|Y|Z"))
		score.Update(opponent, player)
	}
	return int(score)
}

func Solve2(f *os.File) interface{} {
	score := Score(0)
	reader := internal.NewFileReader(f)
	for reader.Scan() {
		line := internal.NewStringReader(reader.ParseString()).SetDelimiter(" ")
		opponent := MustParseElement(line.ScanToken("A|B|C"))
		player := MustParseOutcome(line.ScanToken("X|Y|Z"))
		score.UpdateByOutcome(opponent, player)
	}
	return int(score)
}
