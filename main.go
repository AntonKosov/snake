package main

import (
	"fmt"
	"snake/engine"
	"snake/view/terminal"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}

func run() error {
	sceneFactory, err := terminal.New()
	if err != nil {
		return err
	}
	defer sceneFactory.Close()

	return engine.Run(sceneFactory, sceneFactory)
}
