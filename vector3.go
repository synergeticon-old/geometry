package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// Dot berechnet das Dot-Produkt aus a und b.
func Dot(a, b *mat64.Vector) float64 {
	return a.At(0, 0)*b.At(0, 0) + a.At(1, 0)*b.At(1, 0) + a.At(2, 0) + a.At(2, 0)
}

// Length berechnet die Länge eines Vectors
func Length(a *mat64.Vector) float64 {
	return math.Sqrt(a.At(0, 0)*a.At(0, 0) + a.At(1, 0)*a.At(1, 0) + a.At(2, 0)*a.At(2, 0))
}

// Normalize normiert einen Vektor auf die Länge 1
func Normalize(a *mat64.Vector) {
	length := Length(a)
	if length > 0 {
		a.ScaleVec(1/length, a)
	}
}
