package trie

type Node struct {
	Next         map[rune]*Node
	IsEndOfAWord bool
}

func New() *Node {
	return &Node{
		Next: make(map[rune]*Node),
	}
}

// Insert adds a word to the trie.
func (n *Node) Insert(word string) {
	current := n // Use a new variable for traversal
	for _, r := range word {
		if _, ok := current.Next[r]; !ok {
			current.Next[r] = New()
		}
		current = current.Next[r]
	}
	current.IsEndOfAWord = true
}

// Search checks if a word exists in the trie.
func (n *Node) Search(word string) bool {
	current := n

	for _, r := range word {
		if _, ok := current.Next[r]; ok {
			current = current.Next[r]
		} else {
			return false
		}
	}

	return current.IsEndOfAWord
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (n *Node) StartsWith(prefix string) bool {
	current := n
	for _, r := range prefix {
		if _, ok := current.Next[r]; !ok {
			return false
		}
		current = current.Next[r]
	}

	return true
}
