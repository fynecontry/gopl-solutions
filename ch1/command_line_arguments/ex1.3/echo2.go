// Modified version of Echo2 measuring time to run program.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s\nRuntime: %dns\n", s, time.Since(start).Nanoseconds())
}
