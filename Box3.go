package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// Box3 describes a box in 3d by min and max vector
type Box3 struct {
	Min *mat64.Vector
	Max *mat64.Vector
}

func NewBox3(min, max *mat64.Vector) Box3 {
	b3 := Box3{}
	if min == nil {
		b3.Min = mat64.NewVector(3, []float64{
			math.Inf(-1),
			math.Inf(-1),
			math.Inf(-1),
		})
	}

	if max == nil {
		b3.Min = mat64.NewVector(3, []float64{
			math.Inf(1),
			math.Inf(1),
			math.Inf(1),
		})
	}

	return b3
}

func (b3 *Box3) Empty() {
	b3.Min = mat64.NewVector(3, []float64{
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
	})
	b3.Max = mat64.NewVector(3, []float64{
		math.Inf(-1),
		math.Inf(-1),
		math.Inf(-1),
	})
}

func (b3 *Box3) GetCenter() *mat64.Vector {
	center := mat64.NewVector(3, []float64{0, 0, 0})
	center.AddVec(b3.Min, b3.Max)
	MultiplyScalar(center, 0.5)
	return center
}

func (b3 *Box3) SetFromPoints(points []*mat64.Vector) {
	b3.Empty()
	for _, point := range points {
		b3.ExpandByPoint(point)
	}
}

func (b *Box3) ExpandByPoint(vector *mat64.Vector) {
	b.Min = MinVec(b.Min, vector)
	b.Max = MaxVec(b.Max, vector)
}
