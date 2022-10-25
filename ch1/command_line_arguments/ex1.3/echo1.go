// Modified version of Echo1 measuring time to run program.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%s\nRuntime: %d\n", s, time.Since(start).Nanoseconds())
}
