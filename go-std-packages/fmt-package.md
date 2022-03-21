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

## printf vs sprintf vs fprintf

- printf("format", args) is used to print the data onto the standard output which is often a computer monitor.

- sprintf(char \*, "format", args) is like printf. Instead of displaying the formated string on the standard output i.e. a monitor, it stores the formated data in a string pointed to by the char pointer (the very first parameter). The string location is the only difference between printf and sprint syntax.

- fprintf(FILE \*fp, "format", args) is like printf again. Here, instead of displaying the data on the monitor, or saving it in some string, the formatted data is saved on a file which is pointed to by the file pointer which is used as the first parameter to fprintf. The file pointer is the only addition to the syntax of printf.

