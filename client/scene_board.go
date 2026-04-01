package client

import (
	"context"
	"fmt"
	"io"

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
	streamRecv, err := self.conn.AcceptUniStream(context.Background())
	if err != nil {
		return err
	}

	dataAsBytes, err := io.ReadAll(streamRecv)
	if err != nil {
		return err
	}
	data := string(dataAsBytes)
	fmt.Printf("got data: %v\n", data)

	return nil
}

func (self *SceneBoard) Draw(screen *ebiten.Image) {
}

func (self *SceneBoard) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
