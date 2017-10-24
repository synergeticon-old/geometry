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

// MinVec returns a Vector with the minimal components of two vectors
func MinVec(a, b *mat64.Vector) *mat64.Vector {
	return mat64.NewVector(3, []float64{
		math.Min(a.At(0, 0), b.At(0, 0)),
		math.Min(a.At(1, 0), b.At(1, 0)),
		math.Min(a.At(2, 0), b.At(2, 0)),
	})
}

// MinVec returns a Vector with the minimal components of two vectors
func MaxVec(a, b *mat64.Vector) *mat64.Vector {
	return mat64.NewVector(3, []float64{
		math.Max(a.At(0, 0), b.At(0, 0)),
		math.Max(a.At(1, 0), b.At(1, 0)),
		math.Max(a.At(2, 0), b.At(2, 0)),
	})
}

func MultiplyScalar(a *mat64.Vector, scalar float64) {
	a.SetVec(0, a.At(0, 0)*scalar)
	a.SetVec(1, a.At(1, 0)*scalar)
	a.SetVec(2, a.At(2, 0)*scalar)
}
