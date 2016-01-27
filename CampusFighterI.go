package main

import (
	"flag"
	"fmt"
	"github.com/ITR13/campusFighterI/inputManager"
	"github.com/ITR13/campusFighterI/menus"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"strconv"
	"time"
)

func doesMyCodeWork() {
	startTime := time.Now()
	for i := 0; i != -1; i++ {
		fmt.Print(i)
		fmt.Print(": ")
		timeDiff := time.Now().Second() + 60*time.Now().Minute() - startTime.Second() - startTime.Minute()*60
		fmt.Println(timeDiff)
		time.Sleep(time.Second)
	}
}

func doesMyCodeWork2() {
	startTime := time.Now()
	for i := 0; i != -1; i++ {
		fmt.Print(i)
		fmt.Print(": ")
		timeDiff := time.Now().Second() + 60*time.Now().Minute() - startTime.Second() - startTime.Minute()*60
		fmt.Println(timeDiff)
		sdl.Delay(1000)
	}
}

func main() {
	flag.Parse()
	//	go doesMyCodeWork()
	//	go doesMyCodeWork2()

	var x, y int
	if len(os.Args) < 3 {
		fmt.Println("No size was specified, using defaults")
		x = int(640 * 1.5)
		y = int(480 * 1.5)
	} else {
		var err error
		x, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Could not parse " + os.Args[1] + " to int!")
			return
		}
		y, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Could not parse " + os.Args[2] + " to int!")
			return
		}
	}
	window, err := sdl.CreateWindow("Campus Fighter I", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, x, y, sdl.WINDOWPOS_UNDEFINED)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	//sdl.ShowCursor(0)
	window.UpdateSurface()
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	controlManager := inputManager.GetControlManager()

	updateFunctionsToStartWith := make([]inputManager.Update, 1)
	updateFunctionsToStartWith[0] = menus.CreateMenus(window, surface, renderer, &controlManager)
	inputManager.StateMachine(controlManager.ReadInput(), updateFunctionsToStartWith)

	sdl.Quit()
}
