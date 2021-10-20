# Key-Value Table

A simple utility to make key-value tables.

## Usage

```golang
package main

import (
  "github.com/theverything/kvtable"
)

func main() {
  rows := []kvtable.Row{
    {
      Key: "foo",
      Value: "bar",
    },
    {
      Key: "fizz",
      Value: "buzz",
    },
  }
  fmt.Println(kvtable.Create(rows))
}
```

## Output

```txt
+-------------+
| foo  | bar  |
+-------------+
| fizz | buzz |
+-------------+
```