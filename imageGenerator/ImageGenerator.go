package imageGenerator

import (
	"image"
	"image/color"
	"math"
)

type Pointuint16 struct {
	x, y uint16
}

func Point(pointuint16 Pointuint16) image.Point {
	return image.Point{int(pointuint16.x), int(pointuint16.y)}
}

type Line struct {
	points  []Pointuint16
	rounded []uint16
	jagged  []uint16
	width   uint16
	closed  bool
	ColorGradient
}

type Circle struct {
	point  Pointuint16
	radius uint16
	jagged uint16
	width  uint16
	ColorGradient
}

type ColorGradient interface {
	GetColor(x, y uint16) color.RGBA
	BorderColor() color.RGBA
}

type GeneratorSpecifications struct {
	lines   []Line
	circles []Circle
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func GenerateImage(generatorSpecifications GeneratorSpecifications) *image.RGBA {
	minPoint := image.Point{0, 0}
	maxPoint := minPoint
	for i := 0; i < len(generatorSpecifications.lines); i++ {
		currentLen := len(generatorSpecifications.lines[i].points)
		for j := 1; j <= currentLen; j++ {
			currentPoint := Point(generatorSpecifications.lines[i].points[j%currentLen])
			max := Max(int(generatorSpecifications.lines[i].jagged[j%currentLen]), int(generatorSpecifications.lines[i].jagged[j-1])) + int(generatorSpecifications.lines[i].width)
			min := max / 2
			max -= min

			maxPoint.X = Max(maxPoint.X, currentPoint.X+max)
			maxPoint.Y = Max(maxPoint.Y, currentPoint.Y+max)
			minPoint.X = Min(maxPoint.X, currentPoint.X-min)
			minPoint.Y = Min(maxPoint.Y, currentPoint.Y-min)
		}
	}

	for i := 0; i < len(generatorSpecifications.circles); i++ {
		currentPoint := Point(generatorSpecifications.circles[i].point)
		max := int(generatorSpecifications.circles[i].width + generatorSpecifications.circles[i].jagged)
		min := max / 2
		max += int(generatorSpecifications.circles[i].radius) - min
		min += int(generatorSpecifications.circles[i].radius)

		maxPoint.X = Max(maxPoint.X, currentPoint.X+max)
		maxPoint.Y = Max(maxPoint.Y, currentPoint.Y+max)
		minPoint.X = Min(maxPoint.X, currentPoint.X-min)
		minPoint.Y = Min(maxPoint.Y, currentPoint.Y-min)
	}

	ret := image.NewRGBA(image.Rectangle{minPoint, maxPoint})

	for i := 0; i < len(generatorSpecifications.lines); i++ {
		DrawLine(ret, generatorSpecifications.lines[i])
	}
	for i := 0; i < len(generatorSpecifications.circles); i++ {
		DrawCircle(ret, generatorSpecifications.circles[i])
	}

	return ret
}

func DrawLine(myImage *image.RGBA, line Line) {
	for i := 1; i < len(line.points); i++ {
		DrawBasicLine(myImage, float64(line.points[i-1].x), float64(line.points[i-1].y), float64(line.points[i].x)-float64(line.points[i-1].x),
			float64(line.points[i].y)-float64(line.points[i-1].y), float64(line.width), line.ColorGradient)
	}
}

func DrawBasicLine(myImage *image.RGBA, x_0, y_0, x, y, width float64, colorGradient ColorGradient) {
	div := math.Max(math.Abs(x), math.Abs(y))
	x /= div
	y /= div
	for r := float64(0); r < width; r++ {
		DrawBasicCircle(myImage, x_0, y_0, r, colorGradient, true)
	}
	for t := float64(1); t < div; t++ {
		DrawBasicCircle(myImage, x_0+x*t, y_0+y*t, width, colorGradient, true)
	}
	DrawBasicCircle(myImage, x_0+x*div, y_0+y*div, width, colorGradient, true)
}

func DrawCircle(myImage *image.RGBA, circle Circle) {
	for r := float64(0); r < float64(circle.radius)+float64(circle.width/2); r++ {
		if r < float64(circle.radius+circle.width/2) {
			DrawBasicCircle(myImage, float64(circle.point.x), float64(circle.point.y), r, circle.ColorGradient, false)
		} else {
			DrawBasicCircle(myImage, float64(circle.point.x), float64(circle.point.y), r, circle.ColorGradient, true)
		}
	}
}

//Taken from https://en.wikipedia.org/wiki/Midpoint_circle_algorithm
func DrawBasicCircle(myImage *image.RGBA, x_0, y_0, radius float64, colorGradient ColorGradient, isEdge bool) {
	if radius < 0 {
		return
	}
	x := radius
	y := float64(0)
	decisionOver2 := 1 - x
	for y <= x {
		for xM := float64(-1); xM < 2; xM += 2 {
			for yM := float64(-1); yM < 2; yM += 2 {
				if isEdge {
					myImage.SetRGBA(int(x_0+x*xM), int(y_0+y*yM), colorGradient.BorderColor())
					myImage.SetRGBA(int(x_0+y*xM), int(y_0+x*yM), colorGradient.BorderColor())
				} else {
					myImage.SetRGBA(int(x_0+x*xM), int(y_0+y*yM), colorGradient.GetColor(uint16(x_0+x*xM), uint16(y_0+y*yM)))
					myImage.SetRGBA(int(x_0+y*xM), int(y_0+x*yM), colorGradient.GetColor(uint16(x_0+y*xM), uint16(y_0+x*yM)))
				}
			}
		}
		y++
		if decisionOver2 <= 0 {
			decisionOver2 += 2*y + 1
		} else {
			x--
			decisionOver2 += 2*(y-x) + 1
		}
	}
}
