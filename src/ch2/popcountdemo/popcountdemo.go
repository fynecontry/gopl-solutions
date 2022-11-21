// popcountdemo demonstrates popcount package implementations
package main

import (
	"fmt"

	"popcount"
)

func main() {
	fmt.Printf("Popcount of 0x1234567890ABCDEF using Popcount package: %v\n", popcount.PopCount(0x1234567890ABCDEF))
}
