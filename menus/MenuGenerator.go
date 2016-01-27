package menus

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	//	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	LeftRight = 1 << iota
	UpDown    = 1 << iota
	YX        = 1 << iota
	RightLeft = 0
	DownUp    = 0
	XY        = 0
)

type AnimationDirection int

type MenuItemAnimation struct {
	dx, dy                  int32
	highlightedAnimation    *sdl.Surface
	notHighlightedAnimation *sdl.Surface
	FrameLength             uint64
	AnimationDirection      AnimationDirection
}

func (animation *MenuItemAnimation) Generate(renderer *sdl.Renderer) (*sdl.Texture, *sdl.Texture, []*sdl.Rect, uint64) {
	highlightedTexture, err := renderer.CreateTextureFromSurface(animation.highlightedAnimation)
	if err != nil {
		panic(err)
	}
	notHighlightedTexture, err := renderer.CreateTextureFromSurface(animation.notHighlightedAnimation)
	rows := animation.highlightedAnimation.W / animation.dx
	columns := animation.highlightedAnimation.H / animation.dy
	rects := make([]*sdl.Rect, rows*columns)
	xMult := columns * int32((animation.AnimationDirection>>YX)%2)
	if xMult == 0 {
		xMult = 1
	}
	yMult := rows * int32(1-(animation.AnimationDirection>>YX)%2)
	if yMult == 0 {
		yMult = 1
	}
	xDir := int32(1)
	xPluss := int32(0)
	yDir := int32(1)
	yPluss := int32(0)
	if (animation.AnimationDirection>>LeftRight)%2 == 0 {
		xDir = -1
		xPluss = rows
	}
	if (animation.AnimationDirection>>UpDown)%2 == 0 {
		yDir = -1
		yPluss = columns
	}
	for y := int32(0); y < columns; y++ {
		for x := int32(0); x < rows; x++ {
			currentRect := sdl.Rect{animation.dx * x, animation.dy * y, animation.dx, animation.dy}
			rects[((xPluss+x*xDir)*xMult + (yPluss+y*yDir)*yMult)] = &currentRect
		}
	}
	return highlightedTexture, notHighlightedTexture, rects, animation.FrameLength
}

func LoadMenuAnimation(highlightedPath, notHighlightedPath string, dx, dy int32) MenuItemAnimation {
	highlightedAnimation, err := img.Load(highlightedPath)
	if err != nil {
		panic(err)
	}
	notHighlightedAnimation, err := img.Load(notHighlightedPath)
	if err != nil {
		panic(err)
	}
	return MenuItemAnimation{dx, dy, highlightedAnimation, notHighlightedAnimation, 1, LeftRight | UpDown | YX}
}

type MenuGeneratorInfo struct {
	menuBackground    *sdl.Surface
	menuItemAnimation *MenuItemAnimation
	buttonTexts       []string
	//font
}

func LoadMenuGeneratorInfo(menuItemAnimation *MenuItemAnimation, backgroundPath string, buttonTexts []string, fontPath string) MenuGeneratorInfo {
	background, err := img.Load(backgroundPath)
	if err != nil {
		panic(err)
	}
	return MenuGeneratorInfo{background, menuItemAnimation, buttonTexts}
}

//TODO Make sdl_tff work!
func (menuGeneratorInfo *MenuGeneratorInfo) GenerateMenuShellOneColumn(renderer *sdl.Renderer, x, y, dx, dy int32) Menu {
	background, err := renderer.CreateTextureFromSurface(menuGeneratorInfo.menuBackground)
	if err != nil {
		panic(err)
	}

	buttonCount := len(menuGeneratorInfo.buttonTexts)

	menuItems := make([]*MenuItem, buttonCount)
	highlightedBackground, notHighlightedBackground, srcRects, FrameLength := menuGeneratorInfo.menuItemAnimation.Generate(renderer)

	TEXTSURFACE := sdl.Texture{}
	TEXTRECT := sdl.Rect{0, 0, 0, 0}

	for i := 0; i < buttonCount; i++ {
		currentRect := sdl.Rect{x, y + dy*int32(i), dx, dy}
		currentRects := make([]*sdl.Rect, len(srcRects))
		for j := 0; j < len(srcRects); j++ {
			currentRects[j] = &currentRect
		}
		previous := (i - 1 + buttonCount) % buttonCount
		next := (i + 1) % buttonCount
		menuAction := MenuAction{nil, func(menuInfo *MenuInfo, menuItem *MenuItem) int { return previous }, func(menuInfo *MenuInfo, menuItem *MenuItem) int { return next }, nil, nil}
		menuItems[i] = &MenuItem{notHighlightedBackground, highlightedBackground, srcRects, currentRects, &TEXTSURFACE, &TEXTRECT, menuAction, FrameLength}
	}
	return Menu{background, menuItems, nil}
}
