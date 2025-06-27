package main

func lengthOfLongestSubstringAlt(s string) int {
	current := newQ()
	previous := newQ()

	for _, c := range s {
		item := string(c)
		if !contains(current, item) {
			current = add(current, item)

			continue
		}

		if len(current.q) > len(previous.q) {
			previous = current
		}

		current = remove(current, item)
		current = add(current, item)
	}

	if len(current.q) > len(previous.q) {
		return len(current.q)
	} else {
		return len(previous.q)
	}
}

type Queue struct {
	q []string
	m map[string]int
}

func newQ() Queue {
	return Queue{
		q: make([]string, 0),
		m: make(map[string]int),
	}
}

func contains(q Queue, s string) bool {
	_, ok := q.m[s]

	return ok
}

func add(q Queue, s string) Queue {
	q.q = append(q.q, s)
	q.m[s] = len(q.q)

	return q
}

func remove(q Queue, s string) Queue {
	for i := 0; i < len(q.q); i++ {
		if q.q[i] != s {
			continue
		}

		q.q = q.q[i+1:]
		break
	}

	return q
}
