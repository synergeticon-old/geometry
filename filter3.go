package geometry

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

// PassThroughFilter filters Pointclouds by upper and lower limit
type PassThroughFilter struct {
	LimitLow  float64
	LimitHigh float64
	Normal    *mat64.Vector
}

// SetFilterFieldName sets the Normal Vector by axis-name
func (ptf *PassThroughFilter) SetFilterFieldName(field string) {
	switch field {
	case "x":
		ptf.Normal = mat64.NewVector(3, []float64{1, 0, 0})
	case "y":
		ptf.Normal = mat64.NewVector(3, []float64{0, 1, 0})
	case "z":
		ptf.Normal = mat64.NewVector(3, []float64{0, 0, 1})
	default:
		panic("must be x, y or z")
	}
}

// SetCenterLimit sets the lower and upper limits by center and +/- limit
func (ptf *PassThroughFilter) SetCenterLimit(center, limit float64) {
	ptf.LimitLow = center - limit
	ptf.LimitHigh = center + limit
}

// Filter returns a filtered PointCloud of a Input PoinCloud
func (ptf *PassThroughFilter) Filter(pc PointCloud) PointCloud {
	filteredCloud := PointCloud{}

	plane := NewPlane3(ptf.Normal, -ptf.LimitLow)
	limit := ptf.LimitHigh - ptf.LimitLow
	for _, point := range pc.Vectors {
		value := plane.DistanceToPoint(point)
		fmt.Println(value)
		if value < limit && value > 0 {
			filteredCloud.Vectors = append(filteredCloud.Vectors, point)
		}
	}

	return filteredCloud
}
