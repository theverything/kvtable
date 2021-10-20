package kvtable_test

import (
	"strings"
	"testing"

	"github.com/theverything/kvtable"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		want string
		arg  []kvtable.Row
	}{
		{
			name: "single row",
			want: strings.TrimSpace(`
+-----------+
| foo | bar |
+-----------+
			`),
			arg: []kvtable.Row{
				{
					Key:   "foo",
					Value: "bar",
				},
			},
		},
		{
			name: "multiple rows",
			want: strings.TrimSpace(`
+-------------+
| foo  | bar  |
+-------------+
| fizz | buzz |
+-------------+
			`),
			arg: []kvtable.Row{
				{
					Key:   "foo",
					Value: "bar",
				},
				{
					Key:   "fizz",
					Value: "buzz",
				},
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
			arg:  []kvtable.Row{},
		},
		{
			name: "empty row",
			want: strings.TrimSpace(`
+-----+
|  |  |
+-----+
			`),
			arg: []kvtable.Row{
				{
					Key:   "",
					Value: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kvtable.Create(tt.arg); got != tt.want {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
