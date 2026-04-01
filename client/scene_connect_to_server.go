package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quic-go/quic-go"
)

type SceneConnectToServer struct {
	result chan ConnAnderror
}

func NewSceneConnectToServer() *SceneConnectToServer {
	result := make(chan ConnAnderror)

	go connectToServer(result)

	return &SceneConnectToServer{
		result: result,
	}
}

func (self *SceneConnectToServer) Update() error {
	select {
	case result := <-self.result:
		if result.err != nil {
			return result.err
		} else {
			fmt.Printf("got conn: %v\n", result.conn)
			return NewChangeScene(NewSceneBoard(result.conn))
		}
	default:
	}

	return nil
}

func (self *SceneConnectToServer) Draw(screen *ebiten.Image) {
}

func (self *SceneConnectToServer) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type ConnAnderror struct {
	conn *quic.Conn
	err  error
}

func connectToServer(result chan ConnAnderror) {
	defer close(result)

	conn, err := quic.DialAddr(
		context.Background(),
		ServerAddr,
		&tls.Config{
			InsecureSkipVerify: true,
			NextProtos:         []string{QuicProto},
		},
		nil,
	)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	result <- ConnAnderror{
		conn: conn,
		err:  err,
	}
}
