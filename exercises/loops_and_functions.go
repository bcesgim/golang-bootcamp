/*
Exercise: Loops and Functions

As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := 1.0
z := float64(1)
Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	temp := z
	delta := math.Abs(z - temp)
	for i := 1; i < 11; i++ {
		// Initial guess from Newton's method
		z = z - ((z*z - x) / (2 * z))
		tempValue := fmt.Sprintf(" for %vth iteration, sqrt(x) = %v\n", i, z)
		fmt.Println(tempValue)
		// Update this value as long as some condition upon a small delta holds
		switch {
		case delta > 0.000001:
			temp = z
		}
	}
	return z
}

func main() {
	x := 9000.0 // Replace by test value for computing sqrt(x)
	fmt.Printf("El resultado calculado es: %v\n", Sqrt(x))
	fmt.Printf("El resultado real es: %v\n", math.Sqrt(x))
	// fmt.Printf("La diferencia entre ambos es: %v\n", math.Abs(math.Sqrt(x)-Sqrt(x)))
}

/* Comentarios
1. Podría haber encapsulado la lógica del método de Newton en otra función e imprimir los resultados intermedios desde ahí y no desde Sqrt().
2. Implementando 1 podría imprimir la diferencia entre el valor real (usando math) y el calculado sin volver a imprimir los resultados intermedios
*/
