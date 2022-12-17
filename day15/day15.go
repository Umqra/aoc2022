package day15

import (
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Point struct{ x, y int }
type Measurement struct{ sensor, beacon Point }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (a Point) distTo(b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type Beacons map[Point]struct{}

func read(f *os.File) ([]Measurement, Beacons) {
	reader := internal.NewFileReader(f)
	measurements := make([]Measurement, 0)
	beacons := make(Beacons)
	for reader.Scan() {
		line := reader.ParseString()
		var sensor, beacon Point
		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		measurements = append(measurements, Measurement{sensor: sensor, beacon: beacon})
		beacons[beacon] = struct{}{}
	}
	return measurements, beacons
}

func Solve1(f *os.File, y int) interface{} {
	ms, bs := read(f)
	var minX, maxX int
	for _, m := range ms {
		leftX := m.sensor.x - m.sensor.distTo(m.beacon)
		rightX := m.sensor.x + m.sensor.distTo(m.beacon)
		if leftX < minX {
			minX = leftX
		}
		if rightX > maxX {
			maxX = rightX
		}
	}
	cnt := 0
	for x := minX; x <= maxX; x++ {
		ok := true
		for _, m := range ms {
			point := Point{x: x, y: y}
			if _, exist := bs[point]; !exist && m.sensor.distTo(point) <= m.sensor.distTo(m.beacon) {
				ok = false
				break
			}
		}
		if !ok {
			cnt += 1
		}
	}
	return cnt
}

func Solve2(f *os.File, bound int) interface{} {
	ms, bs := read(f)
	for _, m := range ms {
		d := m.sensor.distTo(m.beacon) + 1
		points := []Point{
			{x: m.sensor.x + d, y: m.sensor.y},
			{x: m.sensor.x, y: m.sensor.y + d},
			{x: m.sensor.x - d, y: m.sensor.y},
			{x: m.sensor.x, y: m.sensor.y - d},
		}
		directions := []Point{
			{x: -1, y: +1},
			{x: -1, y: -1},
			{x: +1, y: -1},
			{x: +1, y: +1},
		}
		for i := 0; i < 4; i++ {
			for l := 0; l < d; l++ {
				current := Point{x: points[i].x + l*directions[i].x, y: points[i].y + l*directions[i].y}
				if current.x < 0 || current.x > bound || current.y < 0 || current.y > bound {
					continue
				}
				if _, exist := bs[current]; exist {
					continue
				}

				ok := true
				for _, t := range ms {
					if t.sensor.distTo(t.beacon) >= t.sensor.distTo(current) {
						ok = false
						break
					}
				}
				if ok {
					return int64(current.x)*4000000 + int64(current.y)
				}
			}
		}
	}
	return 0
}
