package main

import "fmt"

func solve(s string, t string) bool {
	var s_map [256]int
	var t_map [256]int

	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]

		if s_map[sc] != t_map[tc] {
			return false
		}

		s_map[sc] = i + 1
		t_map[tc] = i + 1
	}

	return true
}

func resolve(s string, t string) bool {
	sm := make(map[byte]byte)
	tm := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]

		if val, exists := sm[sChar]; exists {
			if val != tChar {
				return false
			}
		} else {
			sm[sChar] = tChar
		}

		if val, exists := tm[tChar]; exists {
			if val != sChar {
				return false
			}
		} else {
			tm[tChar] = sChar
		}
	}

	return true
}

func main() {
	fmt.Println(solve("paper", "title"))
	fmt.Println(resolve("paper", "title"))
}
