package main

import (
	"context"
	"log"
	"snake/engine"
	"snake/view/terminal"
)

func main() {
	errCh := make(chan error, 1)
	sceneFactory, err := terminal.New(errCh)
	if err != nil {
		log.Fatal(err)
	}
	defer sceneFactory.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case err := <-errCh:
			log.Fatal(err)
		case <-ctx.Done():
			return
		}
	}()

	if err := engine.Run(ctx, sceneFactory, sceneFactory); err != nil {
		log.Fatal(err)
	}
}
