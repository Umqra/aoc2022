package day2

import (
	"fmt"
	"os"
	"strings"
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

func ParseElement(s string) (element Element, err error) {
	switch s {
	case "A", "X":
		element = Rock
	case "B", "Y":
		element = Paper
	case "C", "Z":
		element = Scissors
	default:
		err = fmt.Errorf("unable to parse element string: %v", s)
	}
	return
}

func ParseOutcome(s string) (outcome Outcome, err error) {
	switch s {
	case "X":
		outcome = Lose
	case "Y":
		outcome = Draw
	case "Z":
		outcome = Win
	default:
		err = fmt.Errorf("unable to parse outcome string: %v", s)
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

func Solve1(input string) int {
	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(fmt.Errorf("unable to read input: %w", err))
	}
	lines := strings.Split(string(bytes), "\n")
	score := Score(0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic(fmt.Errorf("each line must have exactly 2 tokens"))
		}
		opponent, err := ParseElement(tokens[0])
		if err != nil {
			panic(fmt.Errorf("invalid opponent token: %w", err))
		}
		player, err := ParseElement(tokens[1])
		if err != nil {
			panic(fmt.Errorf("invalid player token: %w", err))
		}
		score.Update(opponent, player)
	}
	return int(score)
}

func Solve2(input string) int {
	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(fmt.Errorf("unable to read input: %w", err))
	}
	lines := strings.Split(string(bytes), "\n")
	score := Score(0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic(fmt.Errorf("each line must have exactly 2 tokens"))
		}
		opponent, err := ParseElement(tokens[0])
		if err != nil {
			panic(fmt.Errorf("invalid opponent token: %w", err))
		}
		outcome, err := ParseOutcome(tokens[1])
		if err != nil {
			panic(fmt.Errorf("invalid outcome token: %w", err))
		}
		score.UpdateByOutcome(opponent, outcome)
	}
	return int(score)
}
