
## understand the io package

1. the writer 

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- type Writer is the interface that wraps the basic Write method.
- Write writes len(p) bytes from p to the underlying data stream. 
- It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.

2. writeString func

```go
func WriteString(w Writer, s string) (n int, err error)
```

WriteString writes the contents of the string s to w, which accepts a slice of bytes. 
If w implements StringWriter, its WriteString method is invoked directly. Otherwise, w.Write is called exactly once.

```go
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}

}
```
