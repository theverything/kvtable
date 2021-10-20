package kvtable

import (
	"fmt"
	"strings"
)

type Row struct {
	Key   string
	Value string
}

func Create(data []Row) string {
	if len(data) == 0 {
		return ""
	}

	extraChars := 5
	var keyPadding int
	var valPadding int
	for _, row := range data {
		if len(row.Key) > keyPadding {
			keyPadding = len(row.Key)
		}

		if len(row.Value) > valPadding {
			valPadding = len(row.Value)
		}
	}

	line := fmt.Sprintf("+%s+", strings.Repeat("-", extraChars+keyPadding+valPadding))
	var table strings.Builder

	fmt.Fprint(&table, line)

	for _, row := range data {
		fmt.Fprintf(&table, "\n| %-*s | %-*s |\n%s", keyPadding, row.Key, valPadding, row.Value, line)
	}

	return table.String()
}
