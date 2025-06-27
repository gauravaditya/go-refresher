package main

import (
	"maps"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	m := make(map[string]int)
	for _, c := range t {
		m[string(c)]++
	}

	indexes := make([]int, 0)
	valueAt := make(map[int]string)
	for i, r := range s { // find indexes of characters in s that are in t
		if _, ok := m[string(r)]; ok {
			indexes = append(indexes, i)
			valueAt[i] = string(r)
		}
	}

	// no characters in s from t
	if len(indexes) == 0 || len(indexes) < len(t)  {
		return ""
	}

	current := make(map[string]int)
	substr := s
	start, end := 0, len(t)-1

	// create initial sequence of matching characters
	for i := start; i <= end; i++ {
		item := valueAt[indexes[i]]
		current[item]++
	}

	for end < len(indexes) {
		if maps.Equal(m, current) && len(substr) > len(s[indexes[start]:indexes[end]+1]) {
			substr = s[indexes[start] : indexes[end]+1]
		}

		current[valueAt[indexes[start]]]--
		start++

		end++
		if end < len(indexes) {
			current[valueAt[indexes[end]]]++
		}
	}

	return substr
}
