package day16

import (
	"bytes"
	"fmt"
	"github.com/Umqra/aoc2022/internal"
	"os"
	"strings"
)

type Tunnel struct {
	id   int
	rate int64
	next []string
}

func read(f *os.File) map[string]Tunnel {
	reader := internal.NewFileReader(f)
	tunnels := make(map[string]Tunnel)
	id := 0
	for reader.Scan() {
		lineReader := bytes.NewReader([]byte(reader.ParseString()))
		var source string
		var rate int64
		var tunnelsWord, leadsWord, valvesWord string
		var targets []string
		_, _ = fmt.Fscanf(lineReader, "Valve %s has flow rate=%d; %s %s to %s ", &source, &rate, &tunnelsWord, &leadsWord, &valvesWord)
		for {
			var target string
			n, err := fmt.Fscanf(lineReader, "%s", &target)
			if n == 0 || err != nil {
				break
			}
			targets = append(targets, strings.Trim(target, ","))
		}
		tunnels[source] = Tunnel{
			id:   id,
			rate: rate,
			next: targets,
		}
		id++
	}
	return tunnels
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func brute(v string, time int64, opened map[string]struct{}, d map[[2]string]int64, tunnels map[string]Tunnel) int64 {
	if time <= 0 {
		return 0
	}
	result := int64(0)
	for next, tunnel := range tunnels {
		if tunnel.rate == 0 {
			continue
		}
		if _, open := opened[next]; open {
			continue
		}
		opened[next] = struct{}{}
		distance := d[[2]string{v, next}]
		result = max(result, brute(next, time-distance-1, opened, d, tunnels)+(time-distance-1)*tunnel.rate)
		delete(opened, next)
	}
	return result
}

var cache = make(map[int64]int64)

func bruteCooperative(v string, time int64, opened map[string]struct{}, d map[[2]string]int64, tunnels map[string]Tunnel) int64 {
	if len(opened) == len(tunnels) {
		return 0
	}
	mask := int64(0)
	for name, tunnel := range tunnels {
		if _, ok := opened[name]; ok {
			mask |= 1 << tunnel.id
		}
	}
	result := int64(0)
	if cached, exist := cache[mask]; exist {
		result = cached
	} else {
		result = brute("AA", 26, opened, d, tunnels)
		cache[mask] = result
	}
	if time > 0 {
		for next, tunnel := range tunnels {
			if tunnel.rate == 0 {
				continue
			}
			if _, open := opened[next]; open {
				continue
			}
			distance := d[[2]string{v, next}]
			if time-distance-1 <= 0 {
				continue
			}

			opened[next] = struct{}{}
			result = max(result, bruteCooperative(next, time-distance-1, opened, d, tunnels)+(time-distance-1)*tunnel.rate)
			delete(opened, next)
		}
	}
	return result
}

func calcDistances(tunnels map[string]Tunnel) map[[2]string]int64 {
	d := make(map[[2]string]int64)
	for source, tunnel := range tunnels {
		for _, next := range tunnel.next {
			d[[2]string{source, next}] = 1
		}
	}
	for mid := range tunnels {
		for a := range tunnels {
			for b := range tunnels {
				da, oka := d[[2]string{a, mid}]
				db, okb := d[[2]string{mid, b}]
				dp, okp := d[[2]string{a, b}]
				if oka && okb && okp {
					d[[2]string{a, b}] = min(dp, da+db)
				} else if oka && okb {
					d[[2]string{a, b}] = da + db
				}
			}
		}
	}
	return d
}

func Solve1(f *os.File) interface{} {
	tunnels := read(f)
	d := calcDistances(tunnels)
	return brute("AA", 30, make(map[string]struct{}), d, tunnels)
}

func Solve2(f *os.File) interface{} {
	tunnels := read(f)
	d := calcDistances(tunnels)
	opened := make(map[string]struct{})
	for source, tunnel := range tunnels {
		if tunnel.rate == 0 {
			opened[source] = struct{}{}
		}
	}
	return bruteCooperative("AA", 26, opened, d, tunnels)
}
