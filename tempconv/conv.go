package tempconv

// CToF converts temperature from Celsius to Fahrenheit degrees
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts temperature from Celsius degrees to Kelvins
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts temperature from Kelvins to Celsius degrees
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

//KToF converts temperature from Kelvins to Fahrenheit degrees
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

// CToF converts temperature from Fahrenheit to Celsius degrees
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts temperature from Fahrenheit degrees to Kelvins
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
