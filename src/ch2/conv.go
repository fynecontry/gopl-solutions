package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celcius.
func FToC(f Fahrenheit) Celsius { return Celcius((f - 32) * 5 / 9) }
