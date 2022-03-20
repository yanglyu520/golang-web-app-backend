# fmt package

## fmt.FPrintf()

Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.

```go
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
``

example:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

}
```
