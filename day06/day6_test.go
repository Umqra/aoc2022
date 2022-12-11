package day06

import (
	"github.com/Umqra/aoc2022/internal"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSignal_DetectStartMarker(t *testing.T) {
	if start, _ := Signal("bvwbjplbgvbhsrlpgdmjqwftvncz").DetectStartMarker(4); start != 5 {
		t.Fatalf("StartMarker(bvwbjplbgvbhsrlpgdmjqwftvncz) != 5")
	}
	if start, _ := Signal("nppdvjthqldpwncqszvftbrmjlhg").DetectStartMarker(4); start != 6 {
		t.Fatalf("StartMarker(nppdvjthqldpwncqszvftbrmjlhg) != 5")
	}
	if start, _ := Signal("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg").DetectStartMarker(4); start != 10 {
		t.Fatalf("StartMarker(nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg) != 10")
	}
}

func TestSolve1(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day6.a.in", func(file *os.File) interface{} {
		start, err := Signal(strings.Trim(internal.NewFileReader(file).ScanString(), " \n\r")).DetectStartMarker(4)
		if err != nil {
			log.Panicf("unable to solve day6: %v", err)
		}
		return start
	}))
}

func TestSolve2(t *testing.T) {
	t.Logf("result: %v\n", internal.Eval("day6.a.in", func(file *os.File) interface{} {
		start, err := Signal(strings.Trim(internal.NewFileReader(file).ScanString(), " \n\r")).DetectStartMarker(14)
		if err != nil {
			log.Panicf("unable to solve day6: %v", err)
		}
		return start
	}))
}
