package table

import (
	"fmt"
	"math"
	"strings"
)

func Create(data [][]string) string {
	if len(data) == 0 {
		return ""
	}

	numCol := len(data[0])

	if numCol == 0 {
		return ""
	}

	colPadding := make([]int, len(data[0]))
	for _, row := range data {
		if len(row) != numCol {
			return ""
		}
		for i, col := range row {
			if len(col) > colPadding[i] {
				colPadding[i] = len(col)
			}
		}
	}

	extraChars := int(math.Max(float64(numCol*3-1), 0))
	lineLen := extraChars

	for _, v := range colPadding {
		lineLen = v + lineLen
	}

	line := fmt.Sprintf("+%s+", strings.Repeat("-", lineLen))
	var table strings.Builder

	fmt.Fprint(&table, line)

	rowTemplate := fmt.Sprintf("\n|%s\n%%s", strings.Repeat(" %-*s |", numCol))
	args := make([]interface{}, 0, numCol*2+1)
	for _, row := range data {
		args = args[:0]
		for i := range colPadding {
			args = append(args, colPadding[i], row[i])
		}
		args = append(args, line)
		fmt.Fprintf(&table, rowTemplate, args...)
	}

	return table.String()
}
