package client

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quic-go/quic-go"
)

type SceneBoard struct {
	conn *quic.Conn
}

func NewSceneBoard(conn *quic.Conn) *SceneBoard {
	return &SceneBoard{
		conn: conn,
	}
}

func (self *SceneBoard) Update() error {
	return nil
}

func (self *SceneBoard) Draw(screen *ebiten.Image) {
}

func (self *SceneBoard) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
