package utils

import (
	"snake/view"
	"snake/view/terminal/fonts"
	"strings"

	tcell "github.com/gdamore/tcell/v2"
)

type Emitter struct {
	screen       tcell.Screen
	screenParams view.ScreenParams

	originX int
	originY int
	centerX int
	centerY int

	spacesRow string
}

func NewEmitter(screen tcell.Screen, screenParams view.ScreenParams) Emitter {
	w, h := screen.Size()
	var spacesRowBuilder strings.Builder
	for i := 0; i < screenParams.MinScreenWidth; i++ {
		spacesRowBuilder.WriteRune(' ')
	}
	e := Emitter{
		screen:       screen,
		screenParams: screenParams,
		originX:      (w - screenParams.MinScreenWidth) / 2,
		originY:      (h - screenParams.MinScreenHeight) / 2,
		centerX:      w / 2,
		centerY:      h / 2,
		spacesRow:    spacesRowBuilder.String(),
	}

	return e
}

func (e Emitter) Clear() {
	e.screen.Clear()
}

func (e Emitter) Show() {
	e.screen.Show()
}

func (e Emitter) SetFieldContent(x, y int, style tcell.Style, char rune) {
	xScreen := e.originX + x*2
	yScreen := e.originY + y + e.screenParams.FieldStartY

	e.screen.SetContent(xScreen, yScreen, char, nil, style)
	e.screen.SetContent(xScreen+1, yScreen, char, nil, style)
}

func (e Emitter) EmitCenteredText(style tcell.Style, font fonts.Font, y int, str string) {
	text := fonts.Generate(font, str)
	x := e.getStartCenteredH(text[0])
	for i, line := range text {
		e.emitStr(x, e.originY+y+i, style, line)
	}
}

func (e Emitter) FillTopRows(style tcell.Style, lines int) {
	for i := 0; i < lines; i++ {
		e.emitStr(e.originX, e.originY+i, style, e.spacesRow)
	}
}

func (e Emitter) getStartCenteredH(str string) int {
	return e.centerX - len(str)/2
}
func (e Emitter) emitStr(x, y int, style tcell.Style, str string) {
	for i, c := range str {
		// fmt.Printf("Character: '%v'\n", c)
		e.screen.SetContent(x+i, y, c, nil, style)
	}
}
