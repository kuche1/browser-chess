package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quic-go/quic-go"
)

type SceneConnectToServer struct {
}

func NewSceneConnectToServer() *SceneConnectToServer {
	return &SceneConnectToServer{}
}

func (self *SceneConnectToServer) Update() error {
	// TODO: this fucking block the "X" button on the window
	conn, err := quic.DialAddr(
		context.Background(),
		ServerAddr,
		&tls.Config{
			InsecureSkipVerify: true,
			NextProtos:         []string{QuicProto},
		},
		nil,
	)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("conn=%v\n", conn)

	return nil
}

func (self *SceneConnectToServer) Draw(screen *ebiten.Image) {
}

func (self *SceneConnectToServer) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
