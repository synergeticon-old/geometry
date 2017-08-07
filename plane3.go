package geometry

import (
	"github.com/gonum/matrix/mat64"
)

// Plane3 Struct
type Plane3 struct {
	Normal   *mat64.Vector
	Constant float64
}

// NewPlane3 creates a new Plane
func NewPlane3(normal *mat64.Vector, constant float64) *Plane3 {
	p := &Plane3{}
	p.Normal = normal
	p.Constant = constant
	return p
}

// Set Plane from Normal and Constant
func (p *Plane3) Set(normal *mat64.Vector, constant float64) {
	p.Normal = normal
	p.Constant = constant
}

// DistanceToPoint calculates the Distance from the Plane to a given Point
func (p *Plane3) DistanceToPoint(point *mat64.Vector) float64 {
	return Dot(p.Normal, point) + p.Constant
}

// SetComponents sets the Planes components
func (p *Plane3) SetComponents(x, y, z, w float64) {
	p.Normal = mat64.NewVector(3, []float64{x, y, z})
	p.Constant = w
}

// SetFromNormalAndCoplanarPoint sets the Plane from normal vector and onepoint containted by the plane
func (p *Plane3) SetFromNormalAndCoplanarPoint(normal *mat64.Vector, point *mat64.Vector) {
	p.Normal = normal
	p.Constant = -Dot(point, p.Normal)
}

// ProjectPoint projects a point on the plane
func (p *Plane3) ProjectPoint(point *mat64.Vector) *mat64.Vector {
	projectedPoint := p.Normal
	scale := -p.DistanceToPoint(point)
	projectedPoint.ScaleVec(scale, projectedPoint)
	projectedPoint.AddVec(projectedPoint, point)
	return projectedPoint
}
