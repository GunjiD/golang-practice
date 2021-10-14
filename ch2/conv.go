package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(C*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
