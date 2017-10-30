package geometry_test

import "testing"
import "github.com/gonum/matrix/mat64"
import "github.com/synergeticon/geometry"
import "fmt"

func TestCatmullRome(t *testing.T) {
	v1 := mat64.NewVector(3, []float64{0, 0, 0})
	v2 := mat64.NewVector(3, []float64{0, 1, 0})
	v3 := mat64.NewVector(3, []float64{0, 2, 0})
	v4 := mat64.NewVector(3, []float64{0, 3, 0})

	pc := geometry.PointCloud{}
	pc.Add(v1, v2, v3, v4)

	cm := geometry.NewCatmullRome3(pc)
	point := cm.GetPoint(0.5)

	// TO-DO: Check Results
	fmt.Println(point)
	t.Error()
}

func TestCurveIntersetctsPlane(t *testing.T) {
	v1 := mat64.NewVector(3, []float64{0, 0, 0})
	v2 := mat64.NewVector(3, []float64{0, 1, 0})
	v3 := mat64.NewVector(3, []float64{0, 2, 0})
	v4 := mat64.NewVector(3, []float64{0, 3, 0})

	pc := geometry.PointCloud{}
	pc.Add(v1, v2, v3, v4)

	cm := geometry.NewCatmullRome3(pc)

	normal := mat64.NewVector(3, []float64{0, 1, 0})
	constant := -2.0
	plane := geometry.NewPlane3(normal, constant)

	intersection, _ := cm.IntersectPlane(plane)

	// TO-DO: Check Results
	fmt.Println(intersection)
	t.Error()
}
