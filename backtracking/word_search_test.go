package backtracking

import "testing"

func Test_wordSearch(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "BFDA",
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "BFDEE",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSearch(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("wordSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
