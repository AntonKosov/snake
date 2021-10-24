package game

import (
	"snake/utils"
)

type direction int

const (
	directionUp direction = iota
	directionRight
	directionLeft
	directionDown
)

var oppositeDirections map[direction]direction
var moves map[direction]utils.Vector

func init() {
	oppositeDirections = map[direction]direction{
		directionUp:    directionDown,
		directionRight: directionLeft,
		directionLeft:  directionRight,
		directionDown:  directionUp,
	}
	moves = map[direction]utils.Vector{
		directionUp:    {X: 0, Y: -1},
		directionRight: {X: 1, Y: 0},
		directionLeft:  {X: -1, Y: 0},
		directionDown:  {X: 0, Y: 1},
	}
}

type snake struct {
	fieldWidth    int
	fieldHeight   int
	body          []utils.Vector // the head is at the end
	headDirection direction
	previousTurn  direction
	freeCellsMap  map[utils.Vector]bool
	freeCells     int
}

func newSnake(fieldHeight, fieldWidth int) *snake {
	s := snake{
		fieldWidth:   fieldWidth,
		fieldHeight:  fieldHeight,
		freeCellsMap: make(map[utils.Vector]bool, fieldHeight*fieldWidth),
	}
	for x := 0; x < fieldWidth; x++ {
		for y := 0; y < fieldHeight; y++ {
			s.setCellOccupation(utils.Vector{X: x, Y: y}, true)
		}
	}
	s.initSnake()
	return &s
}

func (s *snake) setCellOccupation(p utils.Vector, isFree bool) {
	s.freeCellsMap[p] = isFree
	switch isFree {
	case true:
		s.freeCells++
	case false:
		s.freeCells--
	}
}

func (s *snake) initSnake() {
	s.headDirection = directionUp
	s.previousTurn = s.headDirection
	x := s.fieldWidth / 2
	const length = 4
	for i := 0; i < length; i++ {
		y := s.fieldHeight/2 + length - i
		p := utils.Vector{X: x, Y: y}
		s.body = append(s.body, p)
		s.setCellOccupation(p, false)
	}
}

func (s *snake) move() (previousHeadPosition, currentHeadPosition, previousTailPosition utils.Vector, ok bool) {
	offset := moves[s.headDirection]
	previousHeadPosition = s.body[len(s.body)-1]
	currentHeadPosition = previousHeadPosition.Add(offset)
	if isFreeCell, ok := s.freeCellsMap[currentHeadPosition]; !ok || !isFreeCell {
		return utils.Vector{}, utils.Vector{}, utils.Vector{}, false
	}
	previousTailPosition = s.body[0]
	//TODO: does the same slice grows indefinitely?
	s.body = append(s.body[1:], currentHeadPosition)
	s.previousTurn = s.headDirection

	s.setCellOccupation(currentHeadPosition, false)
	s.setCellOccupation(previousTailPosition, true)

	return previousHeadPosition, currentHeadPosition, previousTailPosition, true
}

func (s *snake) increaseTail(p utils.Vector) {
	s.body = append([]utils.Vector{p}, s.body...) //TODO: allocating new memory
	s.setCellOccupation(p, false)
}

func (s *snake) rotateHead(d direction) {
	if s.previousTurn != oppositeDirections[d] {
		s.headDirection = d
	}
}
