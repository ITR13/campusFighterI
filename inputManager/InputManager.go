package inputManager

import (
	"github.com/veandco/go-sdl2/sdl"
)

type OrderAB int

const (
	FirstAreAB  = 1 << iota
	SecondAreAB = 1 << iota
)

type Controller struct {
	Up, Down, Left, Right, A, B       bool
	UpF, DownF, LeftF, RightF, AF, BF bool
	player                            int
	arrowKeys                         [4]sdl.Keycode
	aKey, bKey                        [2]sdl.Keycode
}

func (controller *Controller) SetOrderAB(orderAB OrderAB) {
	if controller.player%2 == 0 {
		if orderAB^SecondAreAB == FirstAreAB {
			controller.aKey[0] = sdl.K_f
			controller.bKey[0] = sdl.K_g
		} else {
			controller.aKey[0] = sdl.K_g
			controller.bKey[0] = sdl.K_f
		}
		if orderAB^FirstAreAB == SecondAreAB {
			controller.aKey[1] = sdl.K_r
			controller.bKey[1] = sdl.K_t
		} else {
			controller.aKey[1] = sdl.K_t
			controller.bKey[1] = sdl.K_r
		}
	} else {
		if orderAB^SecondAreAB == FirstAreAB {
			controller.aKey[0] = sdl.K_k
			controller.bKey[0] = sdl.K_l
		} else {
			controller.aKey[0] = sdl.K_k
			controller.bKey[0] = sdl.K_l
		}
		if orderAB^FirstAreAB == SecondAreAB {
			controller.aKey[1] = sdl.K_i
			controller.bKey[1] = sdl.K_o
		} else {
			controller.aKey[1] = sdl.K_o
			controller.bKey[1] = sdl.K_i
		}
	}
}

func GetController(player int, orderAB OrderAB) Controller {
	var arrowKeys [4]sdl.Keycode
	var aKey, bKey [2]sdl.Keycode
	if player%2 == 0 {
		arrowKeys = [4]sdl.Keycode{sdl.K_w, sdl.K_s, sdl.K_a, sdl.K_d}
		if orderAB^SecondAreAB == FirstAreAB {
			aKey[0] = sdl.K_f
			bKey[0] = sdl.K_g
		} else {
			aKey[0] = sdl.K_g
			bKey[0] = sdl.K_f
		}
		if orderAB^FirstAreAB == SecondAreAB {
			aKey[1] = sdl.K_r
			bKey[1] = sdl.K_t
		} else {
			aKey[1] = sdl.K_t
			bKey[1] = sdl.K_r
		}
	} else {
		arrowKeys = [4]sdl.Keycode{sdl.K_UP, sdl.K_DOWN, sdl.K_LEFT, sdl.K_RIGHT}
		if orderAB^SecondAreAB == FirstAreAB {
			aKey[0] = sdl.K_k
			bKey[0] = sdl.K_l
		} else {
			aKey[0] = sdl.K_k
			bKey[0] = sdl.K_l
		}
		if orderAB^FirstAreAB == SecondAreAB {
			aKey[1] = sdl.K_i
			bKey[1] = sdl.K_o
		} else {
			aKey[1] = sdl.K_o
			bKey[1] = sdl.K_i
		}
	}
	return Controller{false, false, false, false, false, false, false, false, false, false, false, false, player, arrowKeys, aKey, bKey}
}

func (controller *Controller) ResetFrameKey() {
	controller.UpF = false
	controller.DownF = false
	controller.LeftF = false
	controller.RightF = false
	controller.AF = false
	controller.BF = false
}

type ControlManager struct {
	Player1, Player2 Controller
	Running          bool
}

func GetControlManager() ControlManager {
	return ControlManager{GetController(1, FirstAreAB|SecondAreAB), GetController(2, FirstAreAB|SecondAreAB), true}
}

var Running bool

func (controlManager *ControlManager) getDirectPointers() ([16]sdl.Keycode, [16](*bool), [16](*bool)) {
	var keyCodes [16]sdl.Keycode
	keyBools := [16](*bool){&controlManager.Player1.Up, &controlManager.Player1.Down, &controlManager.Player1.Left, &controlManager.Player1.Right, &controlManager.Player1.A, &controlManager.Player1.A, &controlManager.Player1.B, &controlManager.Player1.B, &controlManager.Player2.Up, &controlManager.Player2.Down, &controlManager.Player2.Left, &controlManager.Player2.Right, &controlManager.Player2.A, &controlManager.Player2.A, &controlManager.Player2.B, &controlManager.Player2.B}
	fKeyBools := [16](*bool){&controlManager.Player1.UpF, &controlManager.Player1.DownF, &controlManager.Player1.LeftF, &controlManager.Player1.RightF, &controlManager.Player1.AF, &controlManager.Player1.AF, &controlManager.Player1.BF, &controlManager.Player1.BF, &controlManager.Player2.UpF, &controlManager.Player2.DownF, &controlManager.Player2.LeftF, &controlManager.Player2.RightF, &controlManager.Player2.AF, &controlManager.Player2.AF, &controlManager.Player2.BF, &controlManager.Player2.BF}
	for i := 0; i < 4; i++ {
		keyCodes[i] = controlManager.Player1.arrowKeys[i]
		keyCodes[i+8] = controlManager.Player2.arrowKeys[i]
	}
	for i := 0; i < 2; i++ {
		keyCodes[4+i] = controlManager.Player1.aKey[i]
		keyCodes[6+i] = controlManager.Player1.bKey[i]
		keyCodes[12+i] = controlManager.Player2.aKey[i]
		keyCodes[14+i] = controlManager.Player2.bKey[i]
	}

	return keyCodes, keyBools, fKeyBools
}

func (controlManager *ControlManager) ReadInput() Update {
	Running = true
	allSym, allKeys, allFKeys := controlManager.getDirectPointers()

	return func(state int) int {
		if Running {
			controlManager.Player1.ResetFrameKey()
			controlManager.Player2.ResetFrameKey()
			event := sdl.PollEvent()
			for event != nil {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					Running = false
				case *sdl.KeyUpEvent:
					for i := 0; i < 16; i++ {
						if t.Keysym.Sym == allSym[i] {
							*allKeys[i] = false
						}
					}
				case *sdl.KeyDownEvent:
					for i := 0; i < 16; i++ {
						if t.Keysym.Sym == allSym[i] {
							if t.Repeat == 0 {
								*allKeys[i] = true
								*allFKeys[i] = true
							}
						}
					}
				}
				event = sdl.PollEvent()
			}
		} else {
			controlManager.Running = false
			return -1
		}

		return 0
	}
}
