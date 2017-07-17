package geometry

import (
	"github.com/gonum/matrix/mat64"
)

// Plane3 Struct
type Plane3 struct {
	normal   *mat64.Vector
	constant float64
}

// NewPlane3 creates a new Plane
func NewPlane3(normal *mat64.Vector, constant float64) *Plane3 {
	p := &Plane3{}
	p.normal = normal
	p.constant = constant
	return p
}

// Set Plane from Normal and Constant
func (p *Plane3) Set(normal *mat64.Vector, constant float64) {
	p.normal = normal
	p.constant = constant
}

// DistanceToPoint calculates the Distance from the Plane to a given Point
func (p *Plane3) DistanceToPoint(point *mat64.Vector) float64 {
	return Dot(p.normal, point) + p.constant
}

// SetComponents sets the Planes components
func (p *Plane3) SetComponents(x, y, z, w float64) {
	p.normal = mat64.NewVector(3, []float64{x, y, z})
	p.constant = w
}

// SetFromNormalAndCoplanarPoint sets the Plane from normal vector and onepoint containted by the plane
func (p *Plane3) SetFromNormalAndCoplanarPoint(normal *mat64.Vector, point *mat64.Vector) {
	p.normal = normal
	p.constant = -Dot(point, p.normal)
}
