// popcountv1demo demonstrates popcountv1 package implementations.
package main

import (
	"fmt"

	"popcountv1"
)

func main() {
	fmt.Printf("Popcount of 0x1234567890ABCDEF using popcountv1 package: %v\n", popcountv1.PopCount(0x1234567890ABCDEF))
}
