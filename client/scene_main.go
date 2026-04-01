// TODO: unused

package client

import "github.com/hajimehoshi/ebiten/v2"

type SceneMain struct {
}

func NewSceneMain() *SceneMain {
	return &SceneMain{}
}

func (self *SceneMain) Update() error {
	return NewChangeScene(NewSceneConnectToServer())
}

func (self *SceneMain) Draw(screen *ebiten.Image) {
}

func (self *SceneMain) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
