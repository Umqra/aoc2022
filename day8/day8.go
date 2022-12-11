package day8

import (
	"github.com/Umqra/aoc2022/internal"
	"os"
)

type Forest struct {
	height  [][]int
	visible [][]bool
	factor  [][]int
}

func (f Forest) Rotate() Forest {
	h := make([][]int, len(f.height[0]))
	v := make([][]bool, len(f.height[0]))
	k := make([][]int, len(f.height[0]))
	for s := 0; s < len(f.height[0]); s++ {
		h[s] = make([]int, len(f.height))
		v[s] = make([]bool, len(f.height))
		k[s] = make([]int, len(f.height))
		for i := 0; i < len(f.height); i++ {
			h[s][i] = f.height[i][s]
			v[s][i] = f.visible[i][s]
			k[s][i] = f.factor[i][s]
		}
	}
	return Forest{height: h, visible: v, factor: k}
}

type Element struct {
	value    int
	position int
}
type Stack []Element

func (s *Stack) UpdateWith(v Element) int {
	next := *s
	for len(next) > 0 && next[len(next)-1].value < v.value {
		next = next[:len(next)-1]
	}
	next = append(next, v)
	*s = next
	if len(next) == 1 {
		return -1
	}
	return next[len(next)-2].position
}

func (f Forest) ShineHorizontally() {
	for i := 0; i < len(f.height); i++ {
		f.visible[i][0], f.visible[i][len(f.height[i])-1] = true, true
		f.factor[i][0], f.factor[i][len(f.height[i])-1] = 0, 0
		left := Stack(make([]Element, 0))
		left = append(left, Element{value: f.height[i][0], position: 0})
		for s := 1; s < len(f.height[i]); s++ {
			p := left.UpdateWith(Element{value: f.height[i][s], position: s})
			if p == -1 {
				f.visible[i][s] = true
			}
			factor := s
			if p > 0 {
				factor -= p
			}
			f.factor[i][s] *= factor
		}
		right := Stack(make([]Element, 0))
		right = append(right, Element{value: f.height[i][len(f.height[i])-1], position: 0})
		for s := len(f.height[i]) - 2; s >= 0; s-- {
			p := right.UpdateWith(Element{value: f.height[i][s], position: len(f.height[i]) - 1 - s})
			if p == -1 {
				f.visible[i][s] = true
			}
			factor := len(f.height[i]) - 1 - s
			if p > 0 {
				factor -= p
			}
			f.factor[i][s] *= factor
		}
	}
}

func (f Forest) CountVisible() int {
	f.ShineHorizontally()
	r := f.Rotate()
	r.ShineHorizontally()

	total := 0
	for _, row := range r.visible {
		for _, element := range row {
			if element {
				total++
			}
		}
	}
	return total
}

func (f Forest) CalculateScore() int {
	f.ShineHorizontally()
	r := f.Rotate()
	r.ShineHorizontally()

	best := 0
	for _, row := range r.factor {
		for _, element := range row {
			if element > best {
				best = element
			}
		}
	}
	return best
}

func readForest(f *os.File) Forest {
	reader := internal.NewFileReader(f)
	table := make([][]int, 0)
	visible := make([][]bool, 0)
	factor := make([][]int, 0)
	for reader.Scan() {
		line := reader.ParseString()
		row := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = int(line[i] - '0')
		}
		table = append(table, row)
		visible = append(visible, make([]bool, len(line)))
		f := make([]int, len(line))
		for i := 0; i < len(f); i++ {
			f[i] = 1
		}
		factor = append(factor, f)
	}
	return Forest{height: table, visible: visible, factor: factor}

}

func Solve1(f *os.File) interface{} {
	return readForest(f).CountVisible()
}

func Solve2(f *os.File) interface{} {
	return readForest(f).CalculateScore()
}
