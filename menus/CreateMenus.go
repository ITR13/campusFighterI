package menus

import (
	"github.com/ITR13/campusFighterI/inputManager"
	"github.com/ITR13/campusFighterI/testKeys"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

func CreateMenus(window *sdl.Window, surface *sdl.Surface, renderer *sdl.Renderer, controlManager *inputManager.ControlManager) inputManager.Update {
	progress := float64(0)
	mutex := sync.Mutex{}
	dx := 4 * surface.W / 8
	startRect := sdl.Rect{surface.W / 4, 4 * surface.H / 7, 0, 2 * surface.H / 7}
	midBackgroundRect := sdl.Rect{surface.W / 4, 4 * surface.H / 7, dx, 2 * surface.H / 7}
	backgroundRect := sdl.Rect{surface.W/4 - 5, 4*surface.H/7 - 5, dx + 10, 2*surface.H/7 + 10}
	var mainMenu *Menu
	go asyncMenuCreator(renderer, &mutex, &progress, &mainMenu, surface.W, surface.H, window, controlManager)
	return func(state int) int {
		if controlManager.Player1.BF || controlManager.Player2.BF {
			controlManager.Running = false
			return -1
		}
		mutex.Lock()
		startRect.W = int32(float64(dx) * progress)
		surface.FillRect(&backgroundRect, 0xffff0000)
		surface.FillRect(&midBackgroundRect, 0xff000000)
		surface.FillRect(&startRect, 0xffff0000)
		window.UpdateSurface()
		if progress == 1 {
			menuInfo := MenuInfo{0, controlManager, renderer, &sdl.Rect{0, 0, surface.W, surface.H}}
			inputManager.UpdateFunctions = append(inputManager.UpdateFunctions, mainMenu.Open(0, &menuInfo))
			return -1
		}
		mutex.Unlock()
		return 0
	}
}

func asyncMenuCreator(renderer *sdl.Renderer, mutex *sync.Mutex, progress *float64, mainMenuPointer **Menu, W, H int32, window *sdl.Window, controlManager *inputManager.ControlManager) {
	totalItems := float64(4)
	menuItemAnimation := LoadMenuAnimation("./menus/Resources/SampleButtonHighlighted.png", "./menus/Resources/SampleButton.png", 200, 100)
	menuItemAnimation.FrameLength = 60
	mutex.Lock()
	*progress = float64(1) / totalItems
	mutex.Unlock()
	mainMenuGeneratorInfo := LoadMenuGeneratorInfo(&menuItemAnimation, "./menus/Resources/SampleBackground.png", []string{"Random Match", "Custom Match", "Options", "About", "Test Keys"}, "somefont")
	mutex.Lock()
	*progress = float64(2) / totalItems
	mutex.Unlock()
	mainMenu := mainMenuGeneratorInfo.GenerateMenuShellOneColumn(renderer, W/2-100, 50, 200, 100)
	mutex.Lock()
	*progress = float64(3) / totalItems
	mutex.Unlock()
	mainMenu.exitMenu = func(menuInfo *MenuInfo) {
		inputManager.Running = false
	}
	mainMenu.menuItems[4].menuAction.Select = func(menuInfo *MenuInfo, menuItem *MenuItem) int {
		inputManager.UpdateFunctions = append(inputManager.UpdateFunctions, testKeys.StartTestKeys(window, controlManager, mainMenu.Open(4, menuInfo)))
		return -1
	}
	mutex.Lock()
	*progress = float64(1)
	*mainMenuPointer = &mainMenu
	mutex.Unlock()
}
