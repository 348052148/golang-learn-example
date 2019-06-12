package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells	= 100
	xyrange	= 30.0
	xyscale	= width / 2 / xyrange
	zscale 	= height * 0.4
	angle	= math.Pi / 6
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(writer,"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0;i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(writer,"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintf(writer,"</svg>")
	})
	http.ListenAndServe(":8080",nil)
}
func corner(i, j int)(float64,float64)  {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	//z := 1.0//f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale //- z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
func test()  {
	var f float32 = 1 << 24
	fmt.Println(f)
	fmt.Println(f == f + 1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
