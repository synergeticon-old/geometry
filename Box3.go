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
		b3.Max = mat64.NewVector(3, []float64{
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

func (b3 *Box3) SetFromCenterAndSize(center, size *mat64.Vector) {
	v1 := mat64.NewVector(3, []float64{0, 0, 0})
	v1.CloneVec(size)
	MultiplyScalar(v1, 0.5)
	b3.Min.AddScaledVec(center, -0.5, size)
	b3.Max.AddScaledVec(center, 0.5, size)
}

func (b3 *Box3) ExpandByPoint(vector *mat64.Vector) {
	b3.Min = MinVec(b3.Min, vector)
	b3.Max = MaxVec(b3.Max, vector)
}

func (b3 *Box3) ContainsPoint(point *mat64.Vector) bool {
	notContained := point.At(0, 0) < b3.Min.At(0, 0) || point.At(0, 0) > b3.Max.At(0, 0) ||
		point.At(1, 0) < b3.Min.At(1, 0) || point.At(1, 0) > b3.Max.At(1, 0) ||
		point.At(2, 0) < b3.Min.At(2, 0) || point.At(2, 0) > b3.Max.At(2, 0)
	return !notContained
}

func (b3 *Box3) ExpandByScalar(scalar float64) {
	AddScalar(b3.Min, -scalar*0.5)
	AddScalar(b3.Max, scalar*0.5)
}

func (b3 *Box3) GetSize() *mat64.Vector {
	size := mat64.NewVector(3, []float64{0, 0, 0})
	size.SubVec(b3.Max, b3.Min)
	return size
}

func (b3 *Box3) Volume() float64 {
	size := b3.GetSize()
	return size.At(0, 0) * size.At(1, 0) * size.At(2, 0)
}
