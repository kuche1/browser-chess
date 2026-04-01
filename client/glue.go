package client

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Glue struct {
	scene ebiten.Game
}

func NewGlue() *Glue {
	return &Glue{
		scene: NewSceneMain(),
	}
}

func (self *Glue) Update() error {
	var changeScene *ChangeScene

	err := self.scene.Update()
	if err != nil {
		if errors.As(err, &changeScene) {
			self.scene = changeScene.newScene
			// do not call update on the new scene so that we have at least one
			// "empty" update call before attempting to connect to the server
			// so that the window pops up BEFORE having us connect
		} else {
			return err
		}
	}

	return nil
}

func (self *Glue) Draw(screen *ebiten.Image) {
	self.scene.Draw(screen)
}

func (self *Glue) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return self.scene.Layout(outsideWidth, outsideHeight)
}
