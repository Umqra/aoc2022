package day4

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	l, r int
}

func (a Range) Contains(b Range) bool {
	return a.l <= b.l && b.r <= a.r
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (a Range) Overlaps(b Range) bool {
	return Max(a.l, b.l) <= Min(a.r, b.r)
}

func parseRanges(s string) []Range {
	ranges := strings.Split(s, ",")
	result := make([]Range, 0)
	for _, r := range ranges {
		borders := strings.Split(r, "-")
		a, err := strconv.Atoi(borders[0])
		if err != nil {
			panic(fmt.Errorf("unable to parse range %s: %w", r, err))
		}
		b, err := strconv.Atoi(borders[1])
		if err != nil {
			panic(fmt.Errorf("unable to parse range %s: %w", r, err))
		}
		result = append(result, Range{l: a, r: b})
	}
	return result
}

func Solve1(f *os.File) int {
	reader := internal.NewFileReader(f)
	count := 0
	for reader.Scan() {
		line := reader.ParseString()
		ranges := parseRanges(line)
		if ranges[0].Contains(ranges[1]) || ranges[1].Contains(ranges[0]) {
			count++
		}
	}
	return count
}

func Solve2(f *os.File) int {
	reader := internal.NewFileReader(f)
	count := 0
	for reader.Scan() {
		line := reader.ParseString()
		ranges := parseRanges(line)
		if ranges[0].Overlaps(ranges[1]) {
			count++
		}
	}
	return count
}
