package day06

import "fmt"

type Signal string

func (s Signal) DetectStartMarker(l int) (int, error) {
	for i := l - 1; i < len(s); i++ {
		marker := s[i-l+1 : i+1]
		uniq := make(map[rune]struct{})
		for _, c := range marker {
			uniq[c] = struct{}{}
		}
		if len(uniq) == l {
			return i + 1, nil
		}
	}
	return 0, fmt.Errorf("")
}
