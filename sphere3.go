package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// Sphere3 describes a 3D sphere by center and radius
type Sphere3 struct {
	Center *mat64.Vector
	Radius float64
}

// SetFromPoints sets Center and Radius from Points
func (sp *Sphere3) SetFromPoints(points []*mat64.Vector) {
	box := NewBox3(nil, nil)
	box.SetFromPoints(points)

	sp.Center = box.GetCenter()

	var maxRadiusSq float64
	for _, point := range points {
		dist := DistanceSquared(sp.Center, point)
		maxRadiusSq = math.Max(maxRadiusSq, dist)
	}

	sp.Radius = math.Sqrt(maxRadiusSq)

}
