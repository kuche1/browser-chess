package client

import "github.com/hajimehoshi/ebiten/v2"

type ChangeScene struct {
	newScene ebiten.Game
}

func NewChangeScene(newScene ebiten.Game) *ChangeScene {
	return &ChangeScene{
		newScene: newScene,
	}
}

func (self *ChangeScene) Error() string {
	return "ChangeScene"
}
