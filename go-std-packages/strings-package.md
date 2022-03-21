# Strings package

1. func fields(s string) []string

Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space
 
```go
package main

import (
"fmt"
"strings"
)

func main() {
fmt.Printf("Fields are: %q", strings.Fields(" foo bar baz "))
}
```
2. 