package table_test

import (
	"strings"
	"testing"

	"github.com/theverything/table"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		want string
		arg  [][]string
	}{
		{
			name: "single row",
			want: strings.TrimSpace(`
+-----------+
| foo | bar |
+-----------+
			`),
			arg: [][]string{
				{"foo", "bar"},
			},
		},
		{
			name: "multiple rows",
			want: strings.TrimSpace(`
+-----------------------------+
| foo  | bar  | bing  | bong  |
+-----------------------------+
| fizz | buzz | hello | world |
+-----------------------------+
| one  | two  | three | four  |
+-----------------------------+
| 1    | 2    | 3     | 4     |
+-----------------------------+
					`),
			arg: [][]string{
				{"foo", "bar", "bing", "bong"},
				{"fizz", "buzz", "hello", "world"},
				{"one", "two", "three", "four"},
				{"1", "2", "3", "4"},
			},
		},
		{
			name: "multiple with empty values rows",
			want: strings.TrimSpace(`
+----------------------------+
| foo | bar  |       | bong  |
+----------------------------+
|     | buzz | hello | world |
+----------------------------+
| one | two  | three |       |
+----------------------------+
| 1   |      | 3     | 4     |
+----------------------------+
					`),
			arg: [][]string{
				{"foo", "bar", "", "bong"},
				{"", "buzz", "hello", "world"},
				{"one", "two", "three", ""},
				{"1", "", "3", "4"},
			},
		},
		{
			name: "nil arg",
			want: "",
			arg:  nil,
		},
		{
			name: "zero rows",
			want: "",
			arg:  [][]string{},
		},
		{
			name: "empty row",
			want: "",
			arg: [][]string{
				{},
			},
		},
		{
			name: "column miss match",
			want: "",
			arg: [][]string{
				{"foo"},
				{"fizz", "buzz"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := table.Create(tt.arg); got != tt.want {
				t.Errorf("Create()\n%v\nwant\n%v", got, tt.want)
			}
		})
	}
}
