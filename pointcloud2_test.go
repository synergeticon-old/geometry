package geometry

import (
	"image/color"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func TestVecToXYs(t *testing.T) {
	pc := PointCloud{}
	pc.FillRandom(100)

	p, err := plot.New()
	if err != nil {
		t.Error(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", VecToXYs(pc.Vectors, 0, 1))
	if err != nil {
		panic(err)
	}
	ShowPlot(p, 5, 10)

}

func TestShowPlot(t *testing.T) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// A quadratic function x^2
	quad := plotter.NewFunction(func(x float64) float64 { return x * x })
	quad.Color = color.RGBA{B: 255, A: 255}
	p.Add(quad)
	p.Legend.Add("x^2", quad)

	p.X.Min = 0
	p.X.Max = 10
	p.Y.Min = 0
	p.Y.Max = 100

	ShowPlot(p, 5, 10)

}
