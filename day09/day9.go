package day09

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Point struct {
	x, y int
}

type Roadmap struct {
	visited map[Point]struct{}
	rope    []Point
}

func NewRoadmap(length int) Roadmap {
	origin := Point{x: 0, y: 0}
	visited := map[Point]struct{}{}
	visited[origin] = struct{}{}
	rope := make([]Point, length)
	for i := 0; i < length; i++ {
		rope[i] = origin
	}
	return Roadmap{
		visited: visited,
		rope:    rope,
	}
}

var (
	Left  = Point{x: -1, y: 0}
	Right = Point{x: +1, y: 0}
	Down  = Point{x: 0, y: -1}
	Up    = Point{x: 0, y: +1}
)

func (r *Point) Move(direction Point) {
	r.x += direction.x
	r.y += direction.y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return +1
}

func (r *Point) Follow(p Point) {
	dx := Abs(r.x - p.x)
	dy := Abs(r.y - p.y)
	if dx <= 1 && dy <= 1 {
		return
	}
	if dx < dy {
		r.x = p.x
		r.y = p.y + Sign(r.y-p.y)
	} else if dy < dx {
		r.y = p.y
		r.x = p.x + Sign(r.x-p.x)
	} else {
		r.x = p.x + Sign(r.x-p.x)
		r.y = p.y + Sign(r.y-p.y)
	}
}

func (r *Roadmap) Move(direction Point) {
	r.rope[0].Move(direction)
	for i := 1; i < len(r.rope); i++ {
		r.rope[i].Follow(r.rope[i-1])
	}
	r.visited[r.rope[len(r.rope)-1]] = struct{}{}
}

func Trace(f *os.File, roadmap Roadmap) interface{} {
	reader := internal.NewFileReader(f)
	for reader.Scan() {
		lineReader := internal.NewStringReader(reader.ParseString()).SetDelimiter(" ")
		directionSymbol := lineReader.ScanToken("L|R|U|D")
		length := lineReader.ScanInt()
		var direction Point
		switch directionSymbol {
		case "L":
			direction = Left
		case "R":
			direction = Right
		case "U":
			direction = Up
		case "D":
			direction = Down
		default:
			panic(fmt.Sprintf("unexpected direction symbol: %v", directionSymbol))
		}
		for length > 0 {
			roadmap.Move(direction)
			length--
		}
	}
	return len(roadmap.visited)
}

func Solve1(f *os.File) interface{} {
	return Trace(f, NewRoadmap(2))
}

func Solve2(f *os.File) interface{} {
	return Trace(f, NewRoadmap(10))
}
