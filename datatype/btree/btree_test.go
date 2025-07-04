package btree

import (
	"testing"
)

func TestPreOrderIterative(t *testing.T) {
	node := map[int]*Node{
		1:  NewNode(1),
		2:  NewNode(2),
		3:  NewNode(3),
		4:  NewNode(4),
		5:  NewNode(5),
		6:  NewNode(6),
		7:  NewNode(7),
		8:  NewNode(8),
		9:  NewNode(9),
		10: NewNode(10),
		11: NewNode(11),
		12: NewNode(12),
		13: NewNode(13),
		14: NewNode(14),
		15: NewNode(15),
	}

	t.Run("should print 8 -> 4 -> 2 -> 1 -> 3 -> 6 -> 5 -> 7 -> 12 -> 10 -> 9 -> 11 -> 14 -> 13 -> 15", func(t *testing.T) {
		root := node[8].
			appendLeft(
				node[4].
					appendLeft(
						node[2].appendLeft(node[1]).appendRight(node[3]),
					).
					appendRight(
						node[6].appendLeft(node[5]).appendRight(node[7]),
					),
			).
			appendRight(
				node[12].
					appendLeft(
						node[10].appendLeft(node[9]).appendRight(node[11]),
					).
					appendRight(
						node[14].appendLeft(node[13]).appendRight(node[15]),
					),
			)

		PreOrderIterative(root)
	})
}

func TestPostOrderIterative(t *testing.T) {
	node := map[int]*Node{
		1: NewNode(1),
		2: NewNode(2),
		3: NewNode(3),
		4: NewNode(4),
		5: NewNode(5),
	}

	t.Run("should print 4 -> 5 -> 2 -> 3 -> 1", func(t *testing.T) {
		root := node[1].
			appendLeft(
				node[2].
					appendLeft(
						node[4],
					).
					appendRight(
						node[5],
					),
			).
			appendRight(
				node[3],
			)

		PostOrderIterative(root)
	})

}
