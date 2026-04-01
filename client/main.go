package client

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Main() {
	fmt.Printf("client: hi\n")

	sceneBoard := NewSceneBoard()

	// ebiten.SetWindowSize(800, 600)

	ebiten.SetWindowTitle("Chess")

	if err := ebiten.RunGame(sceneBoard); err != nil {
		log.Fatal(err)
	}
}
