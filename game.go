package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SWIDTH    = 1280
	SHEIDGT   = 720
	CELL_SIZE = 10
)

type Game struct {
	sq ant
}

type squareCoords struct {
	x, y int
}

type ant struct {
	x, y, w, h int
	squares    map[squareCoords]bool
	direction  string
}

func (a *ant) directionUpdate() {
	switch a.direction {
	case "TOP":
		a.y -= CELL_SIZE
	case "DOWN":
		a.y += CELL_SIZE
	case "LEFT":
		a.x -= CELL_SIZE
	case "RIGHT":
		a.x += CELL_SIZE
	}
}

func (a *ant) checkCell() bool {
	check := squareCoords{x: a.x, y: a.y}
	_, ok := a.squares[check]
	return ok
}

func (a *ant) turnLeft() {
	switch a.direction {
	case "TOP":
		a.direction = "LEFT"
	case "RIGHT":
		a.direction = "TOP"
	case "DOWN":
		a.direction = "RIGHT"
	case "LEFT":
		a.direction = "DOWN"
	}
}

func (a *ant) turnRight() {
	switch a.direction {
	case "TOP":
		a.direction = "RIGHT"
	case "RIGHT":
		a.direction = "DOWN"
	case "DOWN":
		a.direction = "LEFT"
	case "LEFT":
		a.direction = "TOP"
	}
}

func (g *Game) Update() error {
	fmt.Println(len(g.sq.squares))
	g.sq.directionUpdate()
	if g.sq.checkCell() {
		g.sq.turnLeft()
		delete(g.sq.squares, squareCoords{x: g.sq.x, y: g.sq.y})
	} else {
		g.sq.turnRight()
		g.sq.squares[squareCoords{x: g.sq.x, y: g.sq.y}] = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(g.sq.x), float32(g.sq.y), float32(g.sq.w), float32(g.sq.h), color.White, false)
	for k := range g.sq.squares {
		vector.DrawFilledRect(screen, float32(k.x), float32(k.y), CELL_SIZE, CELL_SIZE, color.White, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SWIDTH, SHEIDGT
}

func main() {
	ebiten.SetWindowSize(SWIDTH, SHEIDGT)
	ebiten.SetWindowTitle("Langton's Ant")
	ebiten.SetTPS(60)

	g := Game{
		sq: ant{x: (SWIDTH / 2) + CELL_SIZE, y: (SHEIDGT / 2) + CELL_SIZE, w: CELL_SIZE, h: CELL_SIZE, squares: make(map[squareCoords]bool)},
	}
	g.sq.direction = "LEFT"

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
