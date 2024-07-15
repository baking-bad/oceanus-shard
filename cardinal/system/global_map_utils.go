package system

import (
	"errors"
)

type Direction int

const (
	Left   Direction = 0
	Top    Direction = 1
	Right  Direction = 2
	Bottom Direction = 3
)

const InitialRowLength = 2
const RowLengthIncrement = 2

func getNextDirection(current Direction) (next Direction) {
	switch current {
	case Left:
		return Top
	case Top:
		return Right
	case Right:
		return Bottom
	case Bottom:
		return Left
	default:
		err := errors.New("unexpected Direction value")
		panic(err)
	}
}

func getNextPointPosition(currentDirection Direction, currentPoint [2]int) (nextPoint [2]int) {
	switch currentDirection {
	case Left:
		return [2]int{currentPoint[0] - 1, currentPoint[1]}
	case Top:
		return [2]int{currentPoint[0], currentPoint[1] + 1}
	case Right:
		return [2]int{currentPoint[0] + 1, currentPoint[1]}
	case Bottom:
		return [2]int{currentPoint[0], currentPoint[1] - 1}
	default:
		err := errors.New("unexpected Direction value")
		panic(err)
	}
}

func getIslandCoordinates(n int) [2]int {
	objectsInRow := InitialRowLength
	direction := Left
	currentIslandsInRow := 0
	points := make([][2]int, 0, n)
	points = append(points, [2]int{0, 0})
	for range n {
		if currentIslandsInRow == objectsInRow {
			direction = getNextDirection(direction)
			if direction == Left || direction == Right {
				objectsInRow += RowLengthIncrement
			}
			currentIslandsInRow = 0
		}
		points = append(points, getNextPointPosition(direction, points[len(points)-1]))
		currentIslandsInRow++
	}
	return points[n]
}
