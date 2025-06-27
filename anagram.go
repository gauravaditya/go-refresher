package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := [128]int{}
	m2 := [128]int{}
	for _, r := range s {
		m1[int(r)] = m1[int(r)] + 1
	}

	for _, r := range t {
		m2[int(r)] = m2[int(r)] + 1
	}

	return m1 == m2
}
