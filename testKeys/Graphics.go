package testKeys

import (
	"github.com/ITR13/campusFighterI/inputManager"
	"github.com/veandco/go-sdl2/sdl"
)

var player1Rects = [7]sdl.Rect{sdl.Rect{64, 0, 64, 64}, sdl.Rect{64, 64, 64, 64}, sdl.Rect{0, 64, 64, 64}, sdl.Rect{128, 64, 64, 64}, sdl.Rect{32, 128, 64, 64}, sdl.Rect{96, 128, 64, 64}}

type ParsedInput [14]bool

func Draw(window *sdl.Window, parsedInput *ParsedInput, loop *bool) inputManager.Update {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(&sdl.Rect{0, 0, surface.W, surface.H}, 0xff000000)

	var player2Rects [6]sdl.Rect
	for i := 0; i < 6; i++ {
		player2Rects[i] = sdl.Rect{(256 + player1Rects[i].X), (player1Rects[i].Y), (player1Rects[i].W), (player1Rects[i].H)}
	}

	return func(state int) int {
		if *loop {
			for i := 0; i < 6; i++ {
				if parsedInput[i] {
					surface.FillRect(&player1Rects[i], 0xffff0000)
				} else {
					surface.FillRect(&player1Rects[i], 0xffffff00)
				}
				if parsedInput[i+6] {
					surface.FillRect(&player2Rects[i], 0xffff0000)
				} else {
					surface.FillRect(&player2Rects[i], 0xffffff00)
				}
			}
			if parsedInput[12] {
				surface.FillRect(&sdl.Rect{192, 160, 64, 64}, 0xffff0000)
			} else {
				surface.FillRect(&sdl.Rect{192, 160, 64, 64}, 0xffffff00)
			}
			if parsedInput[13] {
				surface.FillRect(&sdl.Rect{192, 224, 64, 64}, 0xffff0000)
			} else {
				surface.FillRect(&sdl.Rect{192, 224, 64, 64}, 0xffffff00)
			}
			err := window.UpdateSurface()
			if err != nil {
				panic(err)
			}
		} else {
			return -1
		}
		return 0
	}
}
