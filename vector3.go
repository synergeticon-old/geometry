package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// NewVector3 Creates a new mat64 Vector in 3d
func NewVector3(x, y, z float64) *mat64.Vector {
	return mat64.NewVector(3, []float64{x, y, z})
}

// Dot berechnet das Dot-Produkt aus a und b.
func Dot(a, b *mat64.Vector) float64 {
	return mat64.Dot(a, b)
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

// DistanceSquared returns the squared distance of two given points
func DistanceSquared(a, b *mat64.Vector) float64 {
	dx := a.At(0, 0) - b.At(0, 0)
	dy := a.At(1, 0) - b.At(1, 0)
	dz := a.At(2, 0) - b.At(2, 0)
	return dx*dx + dy*dy + dz*dz
}

// Distance returns the Distance between two vectors
func Distance(a, b *mat64.Vector) float64 {
	distanceSq := DistanceSquared(a, b)
	return math.Sqrt(distanceSq)
}
