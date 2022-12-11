package day01

import (
	"github.com/Umqra/aoc2022/internal"
	"os"
	"sort"
)

func sortGroups(f *os.File) []int {
	inputReader := internal.NewFileReader(f).SetDelimiter("\n\n")
	groups := make([]int, 0)
	for inputReader.Scan() {
		groupReader := internal.NewStringReader(inputReader.ParseString())
		total := 0
		for groupReader.Scan() {
			total += groupReader.ParseInt()
		}
		groups = append(groups, total)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	return groups
}

func Solve1(f *os.File) interface{} {
	groups := sortGroups(f)
	return groups[0]
}

func Solve2(f *os.File) interface{} {
	groups := sortGroups(f)
	return groups[0] + groups[1] + groups[2]
}
