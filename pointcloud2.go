package geometry

import (
	"github.com/gonum/matrix/mat64"

	"gonum.org/v1/plot/plotter"
)

const dpi = 96

// VecToXYs converts an array of mat64.Vector to plotter.XY, by giving x and y.
func VecToXYs(vectors []*mat64.Vector, xIndex, yIndex int) plotter.XYs {
	n := len(vectors)
	pts := make(plotter.XYs, n)

	for i, vec := range vectors {
		pts[i].X = vec.At(xIndex, 0)
		pts[i].Y = vec.At(yIndex, 0)
	}
	return pts
}

// ShowPlot displays a Plot in a Window
// func ShowPlot(p *plot.Plot, h, w int) {
// 	gl.StartDriver(func(driver gxui.Driver) {

// 		m := image.NewRGBA(image.Rect(0, 0, w*dpi, h*dpi))
// 		c := vgimg.NewWith(vgimg.UseImage(m))
// 		p.Draw(draw.New(c))

// 		width, height := w*dpi, h*dpi

// 		theme := dark.CreateTheme(driver)
// 		img := theme.CreateImage()
// 		window := theme.CreateWindow(width, height, "Image viewer")
// 		texture := driver.CreateTexture(m, 1.0)
// 		img.SetTexture(texture)
// 		window.AddChild(img)
// 		window.OnClose(driver.Terminate)
// 	})

// }
