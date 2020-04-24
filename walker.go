package main

import (
  "fmt"
  "math"
  "errors"
  "log"
)

const (
    GO_FORWARDS = 1
    GO_BACKWARDS = 2
    TURN_90_PLUS = 3
    TURN_90_MINUS = 4
)

type Piece struct {
    Position    Coordinate
    Direction   float64
}

type Coordinate struct {
    X int
    Y int
}

type Board struct {
	Length int
	Height int
}

func (c *Coordinate) Update(speed int, direction float64) {
    dY := math.Sin(direction)
    dX := math.Cos(direction)
    c.X += (speed * int(math.Round(dX)))
    c.Y += (speed * int(math.Round(dY)))
}

func (b *Board) MoveTo(c Coordinate) error {
    fmt.Println(c)
    if (c.X >= b.Length || c.X < 0 || c.Y >= b.Height || c.Y < 0) {
        return errors.New("You're in thin air")
    }
    return nil
}

func (p *Piece) Turn(radians float64) {
    p.Direction = math.Remainder((p.Direction + radians), (2 * math.Pi))
}

func (p *Piece) Move(move int, b Board) error {
    switch move {
        case GO_FORWARDS:
            p.Position.Update(1, p.Direction)
        case GO_BACKWARDS:
            p.Position.Update(-1, p.Direction)
        case TURN_90_PLUS:
            p.Turn(math.Pi/2)
        case TURN_90_MINUS:
            p.Turn(-math.Pi/2)
        default:
            log.Fatal("... I've no idea what you mean")
    }

    err := b.MoveTo(p.Position)
    if err != nil {
        return errors.New("You're Dead")
    }

    return nil
}

func Simulate(moves []int, p Piece, b Board) (Coordinate, error) {
    for _, move := range moves {
        err := p.Move(move, b)
        if err != nil {
            return Coordinate{-1, -1}, err
        }
    }
    return p.Position, nil
}

func main() {
    p := Piece{
        Position: Coordinate{2, 2,},
        Direction: 0,
    }

    b := Board{
        Length: 4,
        Height: 4,
    }

    moves := []int{1, 4, 1, 3, 2, 3, 2, 4, 1}

    finalPosition, err := Simulate(moves, p, b)

    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("You survived. Just luck.", finalPosition)
}