/*
Exercise: Errors
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

type ErrNegativeSqrt float64
and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

*/

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	temp := z
	delta := math.Abs(z - temp)

	switch {

	case x < 0:
		{
			return 0, ErrNegativeSqrt(x)
		}

	case x > 0:
		{
			for i := 0; i < 10; i++ {
				z = z - ((z*z - x) / (2 * z))
				switch {
				case delta > 0.000001:
					temp = z
				}
			}
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// Pregunta: El resultado me da
//1.414213562373095 <nil>
//0 cannot Sqrt negativ number: -2
// Cómo podría hacer para que sea solamente
//1.414213562373095
//cannot Sqrt negativ number: -2
