// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //Canvas size in pixel
	cells         = 100                 //Number of grid cells
	xyrange       = 30.0                //Axis range
	xyscale       = width / 2 / xyrange //Pixels per x or y unit
	zscale        = height * .4         //Pixel per z unit
	angle         = math.Pi / 6         //Angle of x, y axes (=30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	// fmt.Printf("Width: %d \nHeight:%d \nCell:%d \nxyrange:%.2f \nxyscale:%.2f \nzscale:%.2f \nangle:%.2f \nsin30:%.2f \ncos30: %.2f\n-----------------\n", width, height, cells, xyrange, xyscale, zscale, angle, sin30, cos30 )
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:green; fill:black; stroke-width: 0.7; background-color:black;' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	//Find point (x,y) at corner cell (i,j)
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	//Compute surface height
	z := f(x, y)

	//Project (x,y,z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //Distance from origin (0,0)
	return (math.Sin(r) / r)
}
