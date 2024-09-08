package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"wizard-maze-game/internal/entities"
	"wizard-maze-game/internal/graphics"
)

type Game struct {
	mouse *entities.Mouse
	maze  *entities.Maze
	tiltX float64
	tiltY float64
}

// NewGame initializes a new game instance
func NewGame() *Game {
	mouse := entities.NewMouse()
	maze := entities.NewMaze()

	return &Game{
		mouse: mouse,
		maze:  maze,
	}
}

// Update handles the game logic updates
func (g *Game) Update() error {
	g.handleInput()                      // Update tilt based on input
	g.mouse.Update(g.tiltX, g.tiltY, g.maze) // Update mouse with maze
	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	graphics.RenderBackground(screen, 800, 600) // Render the background

	// Draw each wall
	for _, wall := range g.maze.Walls {
		graphics.DrawWall(screen, wall.X, wall.Y, wall.Width, wall.Height)
	}

	// Draw each hole
	for _, hole := range g.maze.Holes {
		graphics.DrawHole(screen, hole.X, hole.Y, hole.Radius)
	}

	// Draw the mouse
	g.mouse.Draw(screen)
}

// Layout sets the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

// handleInput processes user input for tilting the table
func (g *Game) handleInput() {
	mouseX, mouseY := ebiten.CursorPosition()

	screenWidth, screenHeight := 800, 600
	g.tiltX = (float64(mouseX) - float64(screenWidth)/2) / float64(screenWidth/2)
	g.tiltY = (float64(mouseY) - float64(screenHeight)/2) / float64(screenHeight/2)

	if g.tiltX > 1 {
		g.tiltX = 1
	} else if g.tiltX < -1 {
		g.tiltX = -1
	}

	if g.tiltY > 1 {
		g.tiltY = 1
	} else if g.tiltY < -1 {
		g.tiltY = -1
	}
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Wizard Maze Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
