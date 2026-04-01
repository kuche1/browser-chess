// TODO: unused

package client

import "github.com/hajimehoshi/ebiten/v2"

type SceneBoard struct {
}

func NewSceneBoard() *SceneBoard {
	return &SceneBoard{}
}

func (self *SceneBoard) Update() error {
	return nil
}

func (self *SceneBoard) Draw(screen *ebiten.Image) {
}

func (self *SceneBoard) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
