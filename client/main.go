package client

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Main() {
	// ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(NewGlue()); err != nil {
		log.Fatal(err)
	}
}
