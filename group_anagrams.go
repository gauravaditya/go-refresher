package main

import (
	"maps"
	"regexp"
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	pattern := regexp.MustCompile("[^a-z]")
	palindromes := make(map[[26]int][]string)

	for _, s := range strs {
		s = pattern.ReplaceAllString(strings.ToLower(s), "")
		key := [26]int{}

		for _, r := range s {
			key[int(r)-97] = key[int(r)-97] + 1
		}

		palindromes[key] = append(palindromes[key], s)
	}

	return slices.Collect(maps.Values(palindromes))
}
