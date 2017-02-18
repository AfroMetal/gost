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
		printInfo("Absolute zero",
			tempconv.AbsoluteZeroC,
			tempconv.CToK(tempconv.AbsoluteZeroC),
			tempconv.CToF(tempconv.AbsoluteZeroC))

		printInfo("Water freezing temperature",
			tempconv.FreezingC,
			tempconv.CToK(tempconv.FreezingC),
			tempconv.CToF(tempconv.FreezingC))

		printInfo("Water boiling temperature",
			tempconv.BoilingC,
			tempconv.CToK(tempconv.BoilingC),
			tempconv.CToF(tempconv.BoilingC))
	}

	for _, t := range flag.Args() {
		t, _ := strconv.ParseFloat(t, 64)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		f := tempconv.Fahrenheit(t)
		format := "%s = %s\n"
		fmt.Printf(format, c, tempconv.CToK(c))
		fmt.Printf(format, c, tempconv.CToF(c))
		fmt.Printf(format, k, tempconv.KToC(k))
		fmt.Printf(format, k, tempconv.KToF(k))
		fmt.Printf(format, f, tempconv.FToC(f))
		fmt.Printf(format, f, tempconv.FToK(f))
	}
}
func printInfo(title string, c tempconv.Celsius, k tempconv.Kelvin, f tempconv.Fahrenheit) {
	fmt.Printf(title+":\n"+
		"\tCelsius:\t%s\n"+
		"\tKelvins:\t%s\n"+
		"\tFahrenheit:\t%s\n",
		c, k, f)
}
