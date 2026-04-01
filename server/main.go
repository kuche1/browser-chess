package server

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

func Main() {
	listener, err := quic.ListenAddr(ServerAddr, generateTLSConfig(), nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for {
		log.Printf("Waiting for new connection...")

		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		go func() {
			log.Printf("Handling connection...")
			err := handleNewConnection(conn)
			if err != nil {
				log.Printf("%v", err)
			}
			log.Printf("Connection handled!")
		}()
	}
}

func handleNewConnection(conn *quic.Conn) error {
	streamSend, err := conn.OpenUniStream()
	if err != nil {
		return err
	}

	streamSend.Write([]byte("fuck you")) // TODO: no error checking
	streamSend.Close()                   // TODO: no error checking

	_, err = conn.AcceptUniStream(context.Background())
	if err != nil {
		return err
	}
	// TODO: read the stream data

	return nil
}
