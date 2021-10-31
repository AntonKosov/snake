package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"snake/engine"
	"snake/view/terminal"
)

func main() {
	var inputErr error
	var engineErr error
	defer func() {
		if inputErr != nil {
			fmt.Printf("Input error: %v\n", inputErr)
			os.Exit(1)
		}
		if engineErr != nil {
			fmt.Printf("Game error: %v\n", engineErr)
			os.Exit(2)
		}
	}()

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
		case inputErr = <-errCh:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()

	engineErr = engine.Run(ctx, sceneFactory, sceneFactory)
}
