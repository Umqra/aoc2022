package day14

import (
	"github.com/Umqra/aoc2022/internal"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type Map map[Point]byte

func sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return +1
	}
	return 0
}

func read(f *os.File) (Map, int) {
	reader := internal.NewFileReader(f)
	m := make(Map)
	height := 0
	for reader.Scan() {
		lineReader := internal.NewStringReader(reader.ParseString()).SetDelimiter(" -> ")
		positions := make([]Point, 0)
		for lineReader.Scan() {
			position := strings.Split(lineReader.ParseString(), ",")
			x, _ := strconv.Atoi(position[0])
			y, _ := strconv.Atoi(position[1])
			if y > height {
				height = y
			}
			positions = append(positions, Point{x: x, y: y})
		}
		for i := 0; i+1 < len(positions); i++ {
			a := positions[i]
			b := positions[i+1]
			for a != b {
				m[a] = '#'
				a.x += sign(b.x - a.x)
				a.y += sign(b.y - a.y)
			}
			m[a] = '#'
		}
	}
	return m, height
}

func fall(m Map, path []Point, ground int) []Point {
	last := path[len(path)-1]
	if last.y > 1000 {
		return nil
	}
	if last.y < ground {
		for _, d := range []int{0, -1, 1} {
			next := Point{x: last.x + d, y: last.y + 1}
			if _, ok := m[next]; !ok {
				return fall(m, append(path, next), ground)
			}
		}
	}
	m[last] = '*'
	return path[:len(path)-1]
}

func Solve1(f *os.File) interface{} {
	m, height := read(f)
	path := []Point{{x: 500, y: 0}}
	cnt := 0
	for {
		path = fall(m, path, 1000+height)
		if path == nil {
			break
		}
		cnt++
	}
	return cnt
}

func Solve2(f *os.File) interface{} {
	m, height := read(f)
	path := []Point{{x: 500, y: 0}}
	cnt := 0
	for {
		path = fall(m, path, height+1)
		cnt++
		if len(path) == 0 {
			break
		}
	}
	return cnt
}
