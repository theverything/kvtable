# Table

A simple utility to make ascii tables.

## Usage

Each row must have the same amount of columns.

```golang
package main

import (
  "github.com/theverything/table"
)

func main() {
  rows := [][]string{
    {"foo", "bar", "bing", "bong"},
    {"fizz", "buzz", "hello", "world"},
    {"one", "two", "three", "four"},
    {"1", "2", "3", "4"},
  }
  fmt.Println(table.Create(rows))
}
```

## Output

```txt
+-----------------------------+
| foo  | bar  | bing  | bong  |
+-----------------------------+
| fizz | buzz | hello | world |
+-----------------------------+
| one  | two  | three | four  |
+-----------------------------+
| 1    | 2    | 3     | 4     |
+-----------------------------+
```