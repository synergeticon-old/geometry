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

	ray.Direction.SubVec(b, a)
	ray.Origin = a

	return ray
}

// DistanceToPoint calculates the shortest Distance from Ray to point
func (ray *Ray) DistanceToPoint(point *mat64.Vector) float64 {

	v1 := mat64.NewVector(3, []float64{0, 0, 0})
	v1.SubVec(point, ray.Origin)

	directionDistance := Dot(v1, ray.Direction)

	// v1.copy( this.direction ).multiplyScalar( directionDistance ).add( this.origin );
	v1.AddScaledVec(ray.Origin, directionDistance, ray.Direction)

	return Distance(point, v1)
}
