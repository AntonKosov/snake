package engine

import "snake/view"

type StartSceneCreator interface {
	CreateStartScene() view.Start
}

type GameSceneCreator interface {
	CreateGameScene() view.Game
}

type GameOverSceneCreator interface {
	CreateGameOverScene() view.GameOver
}

type SceneFactory interface {
	StartSceneCreator
	GameSceneCreator
	GameOverSceneCreator
}
