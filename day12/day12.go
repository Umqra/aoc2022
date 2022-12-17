package day12

import (
	"github.com/Umqra/aoc2022/internal"
	"math"
	"os"
)

type Point struct{ x, y int }

func height(c byte) int {
	if c == 'S' {
		return height('a')
	}
	if c == 'E' {
		return height('z')
	}
	return int(c)
}

type Map []string

func (m Map) at(p Point) int {
	return height(m[p.x][p.y])
}

func FindDistance(m Map, start Point, up bool) map[Point]int {
	d := make(map[Point]int)
	d[start] = 0
	queue := make([]Point, 0)
	queue = append(queue, start)
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx*dy != 0 {
					continue
				}
				next := Point{x: current.x + dx, y: current.y + dy}
				if next.x < 0 || next.y < 0 || next.x >= len(m) || next.y >= len(m[next.x]) {
					continue
				}
				if up && m.at(next) > m.at(current)+1 {
					continue
				}
				if !up && m.at(next) < m.at(current)-1 {
					continue
				}
				if _, ok := d[next]; ok {
					continue
				}
				queue = append(queue, next)
				d[next] = d[current] + 1
			}
		}
	}
	return d
}

func read(f *os.File) (Map, Point, Point) {
	reader := internal.NewFileReader(f)
	m := make(Map, 0)
	var start, end Point
	row := 0
	for reader.Scan() {
		line := reader.ParseString()
		for column := 0; column < len(line); column++ {
			if line[column] == 'S' {
				start = Point{x: row, y: column}
			}
			if line[column] == 'E' {
				end = Point{x: row, y: column}
			}
		}
		m = append(m, line)
		row += 1
	}
	return m, start, end
}

func Solve1(f *os.File) interface{} {
	m, start, end := read(f)
	return FindDistance(m, start, true)[end]
}

func Solve2(f *os.File) interface{} {
	m, _, end := read(f)
	d := FindDistance(m, end, false)
	best := math.MaxInt
	for row := 0; row < len(m); row++ {
		for column := 0; column < len(m[row]); column++ {
			if m.at(Point{x: row, y: column}) != height('a') {
				continue
			}
			if distance, ok := d[Point{x: row, y: column}]; ok && distance < best {
				best = distance
			}
		}
	}
	return best
}
