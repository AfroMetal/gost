// Program temperature converts given number(s) to Celsius, Fahrenheit and Kelvin scale temperatures.
package main

import (
	"flag"
	"fmt"
	"strconv"
	"tempconv"
)

var verbose = flag.Bool("verbose", false, "separator")

func main() {
	flag.Parse()

	if *verbose {
		fmt.Printf("Absolute zero:\n"+
			"\tCelsius:\t%s\n"+
			"\tKelvins:\t%s\n"+
			"\tFahrenheit:\t%s\n",
			tempconv.AbsoluteZeroC,
			tempconv.CToK(tempconv.AbsoluteZeroC),
			tempconv.CToF(tempconv.AbsoluteZeroC))

		fmt.Printf("Water freezing temperature:\n"+
			"\tCelsius:\t%s\n"+
			"\tKelvins:\t%s\n"+
			"\tFahrenheit:\t%s\n",
			tempconv.FreezingC,
			tempconv.CToK(tempconv.FreezingC),
			tempconv.CToF(tempconv.FreezingC))

		fmt.Printf("Water boiling temperature:\n"+
			"\tCelsius:\t%s\n"+
			"\tKelvins:\t%s\n"+
			"\tFahrenheit:\t%s\n\n",
			tempconv.BoilingC,
			tempconv.CToK(tempconv.BoilingC),
			tempconv.CToF(tempconv.BoilingC))
	}

	for _, t := range flag.Args() {
		t, _ := strconv.ParseFloat(t, 64)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		f := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s\n", c, tempconv.CToK(c))
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
		fmt.Printf("%s = %s\n", k, tempconv.KToF(k))
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", f, tempconv.FToK(f))
	}
}
