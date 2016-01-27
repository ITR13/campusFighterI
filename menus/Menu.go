package menus

import (
	"github.com/ITR13/campusFighterI/inputManager"
	"github.com/veandco/go-sdl2/sdl"
)

type MenuInfo struct {
	player         int
	controlManager *inputManager.ControlManager
	renderer       *sdl.Renderer
	screenRect     *sdl.Rect
}

type MenuAction struct {
	Select func(menuInfo *MenuInfo, menuItem *MenuItem) int
	Up     func(menuInfo *MenuInfo, menuItem *MenuItem) int
	Down   func(menuInfo *MenuInfo, menuItem *MenuItem) int
	Left   func(menuInfo *MenuInfo, menuItem *MenuItem) int
	Right  func(menuInfo *MenuInfo, menuItem *MenuItem) int
}

type MenuItem struct {
	notHighlightedBackground *sdl.Texture
	highlightedBackground    *sdl.Texture
	texSrc                   []*sdl.Rect
	texDest                  []*sdl.Rect
	text                     *sdl.Texture
	textSrc                  *sdl.Rect
	menuAction               MenuAction
	frameDiv                 uint64
}

type Menu struct {
	background *sdl.Texture
	menuItems  []*MenuItem
	exitMenu   func(menuInfo *MenuInfo)
}

func (menu *Menu) Open(highlightedItem int, menuInfo *MenuInfo) inputManager.Update {
	renderer := menuInfo.renderer
	running := &menuInfo.controlManager.Running
	subRunning := true

	background := menu.background
	backgroundRect := menuInfo.screenRect

	textureAmount := len(menu.menuItems)
	textures := make([]*sdl.Texture, textureAmount)
	srcRect := make([][]*sdl.Rect, textureAmount)
	destRect := make([][]*sdl.Rect, textureAmount)
	animationFrames := make([]uint64, textureAmount)
	frameDivs := make([]uint64, textureAmount)

	textTextures := make([]*sdl.Texture, textureAmount)
	textSrcRect := make([]*sdl.Rect, textureAmount)
	for i := 0; i < textureAmount; i++ {
		if i == highlightedItem {
			textures[i] = menu.menuItems[i].highlightedBackground
		} else {
			textures[i] = menu.menuItems[i].notHighlightedBackground
		}
		srcRect[i] = menu.menuItems[i].texSrc
		destRect[i] = menu.menuItems[i].texDest
		animationFrames[i] = uint64(len(menu.menuItems[i].texSrc))
		frameDivs[i] = menu.menuItems[i].frameDiv

		textTextures[i] = menu.menuItems[i].text
		textSrcRect[i] = menu.menuItems[i].textSrc
	}
	frame := uint64(0)
	controllers := [2]*inputManager.Controller{&menuInfo.controlManager.Player1, &menuInfo.controlManager.Player2}
	return func(state int) int {
		if state == 0 {
			inputManager.UpdateFunctions = append(inputManager.UpdateFunctions, func(state int) int {
				if *running && subRunning {
					renderer.Clear()
					renderer.Copy(background, backgroundRect, backgroundRect)
					for i := 0; i < textureAmount; i++ {
						currentFrame := (frame / frameDivs[i]) % animationFrames[i]
						renderer.Copy(textures[i], srcRect[i][currentFrame], destRect[i][currentFrame])
						renderer.Copy(textTextures[i], textSrcRect[i], destRect[i][currentFrame])
					}
					frame++
					renderer.Present()
				} else {
					return -1
				}
				return 0
			})
		}
		if *running {
			for i := 0; i < 2; i++ {
				menuInfo.player = i
				if controllers[i].UpF {
					if menu.menuItems[highlightedItem].menuAction.Up != nil {
						next := menu.menuItems[highlightedItem].menuAction.Up(menuInfo, menu.menuItems[highlightedItem])
						if next == -1 {
							subRunning = false
							return -1
						} else if next != highlightedItem {
							textures[highlightedItem] = menu.menuItems[highlightedItem].notHighlightedBackground
							textures[next] = menu.menuItems[highlightedItem].highlightedBackground
							highlightedItem = next
						}
					}
				}
				if controllers[i].DownF {
					if menu.menuItems[highlightedItem].menuAction.Down != nil {
						next := menu.menuItems[highlightedItem].menuAction.Down(menuInfo, menu.menuItems[highlightedItem])
						if next == -1 {
							subRunning = false
							return -1
						} else if next != highlightedItem {
							textures[highlightedItem] = menu.menuItems[highlightedItem].notHighlightedBackground
							textures[next] = menu.menuItems[highlightedItem].highlightedBackground
							highlightedItem = next
						}
					}
				}
				if controllers[i].LeftF {
					if menu.menuItems[highlightedItem].menuAction.Left != nil {
						next := menu.menuItems[highlightedItem].menuAction.Left(menuInfo, menu.menuItems[highlightedItem])
						if next == -1 {
							subRunning = false
							return -1
						} else if next != highlightedItem {
							textures[highlightedItem] = menu.menuItems[highlightedItem].notHighlightedBackground
							textures[next] = menu.menuItems[highlightedItem].highlightedBackground
							highlightedItem = next
						}
					}
				}
				if controllers[i].RightF {
					if menu.menuItems[highlightedItem].menuAction.Right != nil {
						next := menu.menuItems[highlightedItem].menuAction.Right(menuInfo, menu.menuItems[highlightedItem])
						if next == -1 {
							subRunning = false
							return -1
						} else if next != highlightedItem {
							textures[highlightedItem] = menu.menuItems[highlightedItem].notHighlightedBackground
							textures[next] = menu.menuItems[highlightedItem].highlightedBackground
							highlightedItem = next
						}
					}
				}
				if controllers[i].AF {
					if menu.menuItems[highlightedItem].menuAction.Select != nil {
						next := menu.menuItems[highlightedItem].menuAction.Select(menuInfo, menu.menuItems[highlightedItem])
						if next == -1 {
							subRunning = false
							return -1
						} else if next != highlightedItem {
							textures[highlightedItem] = menu.menuItems[highlightedItem].notHighlightedBackground
							textures[next] = menu.menuItems[highlightedItem].highlightedBackground
							highlightedItem = next
						}
					}
				}
				if controllers[i].BF {
					if menu.exitMenu != nil {
						menu.exitMenu(menuInfo)
					}
					subRunning = false
					return -1
				}
			}
		} else {
			return -1
		}
		return 1
	}
}
