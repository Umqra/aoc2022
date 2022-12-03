package day3

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Element byte

type Rucksack struct {
	all    map[Element]struct{}
	first  map[Element]struct{}
	second map[Element]struct{}
}

func parseRucksack(s string) Rucksack {
	if len(s)%2 != 0 {
		panic(fmt.Errorf("rucksack must have even number of items: %v %% 2 != 0", len(s)))
	}
	all, first, second := make(map[Element]struct{}), make(map[Element]struct{}), make(map[Element]struct{})
	for i := 0; i < len(s)/2; i++ {
		all[Element(s[i])] = struct{}{}
		all[Element(s[i+len(s)/2])] = struct{}{}
		first[Element(s[i])] = struct{}{}
		second[Element(s[i+len(s)/2])] = struct{}{}
	}
	return Rucksack{all: all, first: first, second: second}
}

func (rucksack *Rucksack) findCommonElements() []Element {
	common := make([]Element, 0)
	for k, _ := range rucksack.first {
		if _, ok := rucksack.second[k]; ok {
			common = append(common, k)
		}
	}
	return common
}

func findCommonElements(rs []Rucksack) []Element {
	common := make([]Element, 0)
	for k, _ := range rs[0].all {
		ok := true
		for _, r := range rs[1:] {
			if _, ok = r.all[k]; !ok {
				break
			}
		}
		if ok {
			common = append(common, k)
		}
	}
	return common
}

func (e Element) Cost() int {
	if 'a' <= e && e <= 'z' {
		return int(e) - int('a') + 1
	}
	if 'A' <= e && e <= 'Z' {
		return int(e) - int('A') + 27
	}
	panic(fmt.Errorf("unexpected element: %v", e))
}

func Solve1(f *os.File) int {
	reader := internal.NewFileReader(f)
	total := 0
	for reader.Scan() {
		rucksack := parseRucksack(reader.ParseString())
		common := rucksack.findCommonElements()
		for _, element := range common {
			total += element.Cost()
		}
	}
	return total
}

func Solve2(f *os.File) int {
	reader := internal.NewFileReader(f)
	total := 0
	for reader.Scan() {
		r1 := parseRucksack(reader.ParseString())
		r2 := parseRucksack(reader.ScanString())
		r3 := parseRucksack(reader.ScanString())
		common := findCommonElements([]Rucksack{r1, r2, r3})
		for _, element := range common {
			total += element.Cost()
		}
	}
	return total
}
