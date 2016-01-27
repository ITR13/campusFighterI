package character

import "github.com/veandco/go-sdl2/sdl"

type Frame struct {
	Frames      sdl.Rect
	InHitBoxes  []*HitBox
	OutHitBoxes []*HitBox
	HitboxHit   func(inHit, outHit uint64, player *Player)
}

//0 right, 1 left
type Animation struct {
	Texture [2]*sdl.Texture
	Frames  []*Frame
}

type Character struct {
	Idle        Animation
	Crouch      Animation
	Walk        Animation
	Moonwalk    Animation
	Run         Animation
	Jump        Animation
	ForwardJump Animation
	Hurt        Animation
	Recovery    Animation
	Attacks     []Animation
}

type Player struct {
	X, Y, Vx, Vy int
	Character    Character
}
