package main

var isOpeningBrace = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

var isClosingBrace = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isBalancedString(input string) string {
	brackets := make([]rune, 0)

	for _, v := range input {
		if _, ok := isOpeningBrace[v]; ok {
			brackets = append(brackets, v)
			continue
		}

		closing, ok := isClosingBrace[v]
		if !ok {
			continue
		}

		if len(brackets) == 0 {
			continue
		}

		top := brackets[len(brackets)-1]
		if top == closing && len(brackets) > 0 {
			brackets = brackets[:len(brackets)-1]
		}

	}

	if len(brackets) == 0 {
		return "Balanced"
	} else {
		return "Not Balanced"
	}
}
