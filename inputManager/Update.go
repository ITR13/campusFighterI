package inputManager

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Update func(int) int

var UpdateFunctions []Update

func StateMachine(coreFunction Update, updateFunctionsToStartWith []Update) {
	UpdateFunctions = updateFunctionsToStartWith
	functionCount := len(updateFunctionsToStartWith)
	states := make([]int, functionCount)
	for i := 0; i < functionCount; i++ {
		states[i] = 0
	}
	coreState := coreFunction(0)
	for coreState != -1 && functionCount != 0 {
		for i := 0; i < functionCount; i++ {
			states[i] = UpdateFunctions[i](states[i])
		}
		newCount := len(UpdateFunctions)
		states = append(states, make([]int, newCount-functionCount)...)
		for i := functionCount; i < newCount; i++ {
			states[i] = 0
		}
		for i := functionCount - 1; i >= 0; i-- {
			if states[i] == -1 {
				UpdateFunctions = append(UpdateFunctions[:i], UpdateFunctions[i+1:]...)
				states = append(states[:i], states[i+1:]...)
			}
		}
		functionCount = len(UpdateFunctions)
		sdl.Delay(15)
		coreState = coreFunction(coreState)
	}
}
