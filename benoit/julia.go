package benoit

import "math"

// Julia repeats z = 0; z = z^2 + c until z diverges.
// Stops after max repetitions otherwise.
func Julia(c complex128, max int) int {

	var z complex128
	for i := 0; i < max; i++ {
		z = z*z + c

		if mag(z) >= 2 {
			return i
		}
	}

	return max
}

// mag approximates complex number magnitude sqrt(real^2 + imag^2).
func mag(z complex128) float64 {
	return math.Max(math.Abs(real(z)), math.Abs(imag(z)))
}
