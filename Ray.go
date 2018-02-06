package geometry

import (
	"github.com/gonum/matrix/mat64"
)

// Ray represents the Ray object
type Ray struct {
	Origin    *mat64.Vector
	Direction *mat64.Vector
}

// NewRayFromPoints creates a Ray from two points
func NewRayFromPoints(a, b *mat64.Vector) Ray {
	ray := Ray{}

	ray.Direction = mat64.NewVector(3, []float64{0, 0, 0})
	ray.Origin = mat64.NewVector(3, []float64{0, 0, 0})

	ray.Direction.SubVec(b, a)
	ray.Origin.CopyVec(a)

	return ray
}

// DistanceToPoint calculates the shortest Distance from Ray to point
func (ray *Ray) DistanceToPoint(point *mat64.Vector) float64 {

	// Mathematical expression:
	//     |(point-ray.Origin) x ray.Direction|
	// d = ------------------------------------
	//                |ray.Direction|

	v1 := mat64.NewVector(3, []float64{0, 0, 0})

	v1.SubVec(point, ray.Origin)

	// Calculation cross-product of (point-Origin) and the Direction
	v2 := mat64.NewVector(3, []float64{
		(v1.At(1, 0) * ray.Direction.At(2, 0)) - (v1.At(2, 0) * ray.Direction.At(1, 0)),
		(v1.At(2, 0) * ray.Direction.At(0, 0)) - (v1.At(0, 0) * ray.Direction.At(2, 0)),
		(v1.At(0, 0) * ray.Direction.At(1, 0)) - (v1.At(1, 0) * ray.Direction.At(0, 0)),
	})

	// Calculation of Lengths
	len1 := Length(v2)
	len2 := Length(ray.Direction)

	distance := len1 / len2

	return distance
}
