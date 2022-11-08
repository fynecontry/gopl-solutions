// Modified version of Echo3 measuring time to run program.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Runtime: %dns\n", time.Since(start).Nanoseconds())
}
