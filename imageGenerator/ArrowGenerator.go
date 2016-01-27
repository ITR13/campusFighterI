package imageGenerator

import (
	"image/color"
	"math/rand"
)

func GetArrow(seed int64) GeneratorSpecifications {
	randomGen := rand.New(rand.NewSource(seed))
	points := make([]Pointuint16, 12)
	width := 14 + (uint16(randomGen.Int31()) % 6)
	width1 := (uint16(randomGen.Int31()) % 3)
	length1 := 2 + (uint16(randomGen.Int31()) % 6)
	width2 := 3 + (uint16(randomGen.Int31()) % 3)
	length2 := 4 + (uint16(randomGen.Int31()) % 3)
	length3 := (uint16(randomGen.Int31()) % 3)

	points[0] = Pointuint16{0, 0}
	points[7] = Pointuint16{0, width}
	points[1] = Pointuint16{length1, width1}
	points[6] = Pointuint16{length1, width - width1}
	points[2] = Pointuint16{length1 + length2, width1 + width2}
	points[5] = Pointuint16{length1 + length2, width - width1 + width2}
	points[3] = Pointuint16{length1 + length2 + length3, width / 2}
	points[4] = Pointuint16{length1 + length2 + length3, width - width/2}

	length1 = (uint16(randomGen.Int31()) % 6)
	width2 = uint16(randomGen.Int31()) % ((width - 2*width/5) / 2)
	length2 = (uint16(randomGen.Int31()) % 9)

	points[8] = Pointuint16{length1, width / 5}
	points[11] = Pointuint16{length1, width - width/5}
	points[9] = Pointuint16{length1 + length2, width/5 + width2}
	points[10] = Pointuint16{length1 + length2, width - width/5 - width2}

	rounded := (uint16(randomGen.Int31()) % 6)
	rounded2 := (uint16(randomGen.Int31()) % 3)
	rounded3 := (uint16(randomGen.Int31()) % 3)
	roundedData := []uint16{0, rounded, rounded, rounded + rounded2, rounded + rounded2, rounded, rounded, 0, rounded3, rounded3 + rounded2, rounded3 + rounded2, rounded3}

	jagged := (uint16(randomGen.Int31()) % 6)
	jagged2 := (uint16(randomGen.Int31()) % 3)
	jagged3 := (uint16(randomGen.Int31()) % 5)
	jaggedData := []uint16{0, jagged, jagged, jagged + jagged2, jagged + jagged2, jagged, jagged, 0, jagged3, jagged3 + jagged2, jagged3 + jagged2, jagged3}

	lineWidth := (uint16(randomGen.Int31()) % 4)

	arrowFillColor := ArrowFillColor{color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}, color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}, color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}}

	return GeneratorSpecifications{[]Line{Line{points, roundedData, jaggedData, lineWidth, true, arrowFillColor}}, []Circle{}}
}

type ArrowFillColor struct {
	upperLeft, lowerLeft, middleRight color.RGBA
}

func (arrowFillColor ArrowFillColor) GetColor(x, y uint16) color.RGBA {
	return color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
}

func (arrowFillColor ArrowFillColor) BorderColor() color.RGBA {
	return color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
}

/*
	points  []image.Point
	rounded []uint16
	jagged  []uint16
	width   uint16
	closed  bool
	ColorGradient
*/
