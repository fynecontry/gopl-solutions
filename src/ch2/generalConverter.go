// generalConverter displays value passed from CLI or stdin in different unit measurements.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"unitconverter"
)

//var n = flag.Bool("n", false, "parse n values for unit conversion")

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "generalConverter: %v\n", err)
				os.Exit(1)
			}
			displayConversion(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "generalConverter: %v\n", err)
				os.Exit(1)
			}
			displayConversion(t)
		}
	}
}

func displayConversion(t float64) {
	c := unitconverter.Celsius(t)
	ft := unitconverter.Foot(t)
	lb := unitconverter.Pound(t)

	fmt.Printf("Temperature:\t%s = %s\nLength:\t\t%.2f m = %.2f ft\nWeight:\t\t%.2f kg = %.2f lb\n",
		unitconverter.CToF(c), unitconverter.Celsius(t), unitconverter.FToM(ft),
		unitconverter.Foot(ft), unitconverter.LBtoKg(lb), unitconverter.Pound(lb))
}
