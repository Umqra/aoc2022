package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortGroups(input string) []int {
	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(fmt.Errorf("unable to read input: %w", err))
	}
	lines := strings.Split(string(bytes), "\n\n")
	groups := make([]int, len(lines))
	for _, line := range lines {
		total := 0
		group := strings.Split(line, "\n")
		for _, item := range group {
			if item == "" {
				continue
			}
			current, err := strconv.Atoi(item)
			if err != nil {
				panic(fmt.Errorf("unexpected line format: %w", err))
			}
			total += current
		}
		groups = append(groups, total)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	return groups
}

func Solve1(input string) int {
	groups := sortGroups(input)
	return groups[0]
}

func Solve2(input string) int {
	groups := sortGroups(input)
	return groups[0] + groups[1] + groups[2]
}
