package main

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestGroupAnagrams_Basic(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	got := groupAnagrams(input)
	want := [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}
	assertAnagramGroups(t, got, want)
}

func TestGroupAnagrams_EmptyInput(t *testing.T) {
	input := []string{}
	got := groupAnagrams(input)
	want := [][]string{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestGroupAnagrams_SingleWord(t *testing.T) {
	input := []string{"hello"}
	got := groupAnagrams(input)
	want := [][]string{{"hello"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestGroupAnagrams_CaseInsensitiveAndPunctuation(t *testing.T) {
	input := []string{"Eat!", "Tea", "ate", "Ate.", "tEa", "tan", "Nat", "bat"}
	got := groupAnagrams(input)
	want := [][]string{
		{"eat", "tea", "ate", "ate", "tea"},
		{"tan", "nat"},
		{"bat"},
	}
	assertAnagramGroups(t, got, want)
}

func TestGroupAnagrams_AllAnagrams(t *testing.T) {
	input := []string{"abc", "bca", "cab", "cba", "bac", "acb"}
	got := groupAnagrams(input)
	want := [][]string{{"abc", "bca", "cab", "cba", "bac", "acb"}}
	assertAnagramGroups(t, got, want)
}

func TestGroupAnagrams_NoAnagrams(t *testing.T) {
	input := []string{"one", "two", "three"}
	got := groupAnagrams(input)
	want := [][]string{{"one"}, {"two"}, {"three"}}
	assertAnagramGroups(t, got, want)
}

// Helper function to compare groups regardless of order
func assertAnagramGroups(t *testing.T, got, want [][]string) {
	t.Helper()
	normalize := func(groups [][]string) [][]string {
		for _, group := range groups {
			slices.Sort(group)
		}
		slices.SortFunc(groups, func(a, b []string) int {
			if len(a) == 0 || len(b) == 0 {
				return len(a) - len(b)
			}
			return strings.Compare(a[0], b[0])
		})
		return groups
	}
	gotNorm := normalize(got)
	wantNorm := normalize(want)
	if !reflect.DeepEqual(gotNorm, wantNorm) {
		t.Errorf("expected %v, got %v", wantNorm, gotNorm)
	}
}
