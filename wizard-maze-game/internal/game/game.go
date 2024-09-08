package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"wizard-maze-game/internal/entities"
	"wizard-maze-game/internal/graphics"
)

// Screen dimensions
const (
	ScreenWidth  = 800
	ScreenHeight = 600
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
	g.handleInput()
	g.mouse.Update(g.tiltX, g.tiltY)
	g.checkCollisions()
	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Pass the screen width and height to RenderBackground
	graphics.RenderBackground(screen, ScreenWidth, ScreenHeight)
	g.maze.Draw(screen)
	g.mouse.Draw(screen)
	ebitenutil.DebugPrint(screen, "Tilt the table to move the mouse!")
}

// Layout sets the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// handleInput processes user input for tilting the table
func (g *Game) handleInput() {
	x, y := ebiten.CursorPosition()
	g.tiltX = (float64(x) - float64(ScreenWidth)/2) / float64(ScreenWidth/2)
	g.tiltY = (float64(y) - float64(ScreenHeight)/2) / float64(ScreenHeight/2)

	// Clamp tilt values to prevent excessive tilt
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

// checkCollisions checks for collisions with maze walls or holes
func (g *Game) checkCollisions() {
	if entities.CheckCollision(g.mouse, g.maze) {
		// Handle collision logic, e.g., reset game state or reduce lives
	}
}

func main() {
	// Create a new game instance
	game := NewGame()

	// Set up the window size and title
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Wizard Maze Game")

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
