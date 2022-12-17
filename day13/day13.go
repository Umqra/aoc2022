package day13

import (
	"encoding/json"
	"github.com/Umqra/aoc2022/internal"
	"os"
	"sort"
)

type Packet []interface{}

func parse(line string) Packet {
	var result Packet
	err := json.Unmarshal([]byte(line), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func before(a Packet, b Packet) float64 {
	var l int
	if len(a) < len(b) {
		l = len(a)
	} else {
		l = len(b)
	}
	for i := 0; i < l; i++ {
		intA, intACheck := a[i].(float64)
		intB, intBCheck := b[i].(float64)
		if intACheck && intBCheck {
			if intA != intB {
				return intA - intB
			}
			continue
		}
		if !intACheck && !intBCheck {
			if r := before(a[i].([]interface{}), b[i].([]interface{})); r != 0 {
				return r
			}
		}
		if intACheck {
			if r := before(Packet{intA}, b[i].([]interface{})); r != 0 {
				return r
			}
		}
		if intBCheck {
			if r := before(a[i].([]interface{}), Packet{intB}); r != 0 {
				return r
			}
		}
	}
	return float64(len(a) - len(b))
}

func Solve1(f *os.File) interface{} {
	reader := internal.NewFileReader(f)
	i := 1
	total := 0
	for reader.Scan() {
		first := parse(reader.ParseString())
		second := parse(reader.ScanString())
		_ = reader.ScanString()
		if before(first, second) < 0 {
			total += i
		}
		i += 1
	}
	return total
}

func identical(a, b Packet) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return &a[0] == &b[0]
}

func Solve2(f *os.File) interface{} {
	reader := internal.NewFileReader(f)
	packets := make([]Packet, 0)
	divider2 := Packet{[]interface{}{any(float64(2))}}
	packets = append(packets, divider2)
	divider6 := Packet{[]interface{}{any(float64(6))}}
	packets = append(packets, divider6)
	for reader.Scan() {
		first := parse(reader.ParseString())
		second := parse(reader.ScanString())
		_ = reader.ScanString()
		packets = append(packets, first)
		packets = append(packets, second)
	}
	sort.Slice(packets, func(i, j int) bool {
		return before(packets[i], packets[j]) < 0
	})
	result := 1
	for i := 0; i < len(packets); i++ {
		if identical(packets[i], divider2) || identical(packets[i], divider6) {
			result *= i + 1
		}
	}
	return result
}
