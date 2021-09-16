package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		z = z - (z*z-x)/(2*z)
		if z*z-x <= 0 {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(0.0001))
}
