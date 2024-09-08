package entities

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"wizard-maze-game/internal/graphics" // Import the graphics package to access images
)

var (
	// Define the font face globally to avoid reloading it every frame
	fontFace font.Face
)

func init() {
	// Use a basic font for demonstration
	fontFace = basicfont.Face7x13
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

// Mouse represents the player-controlled mouse character
type Mouse struct {
	X, Y          float64   // Position of the mouse
	Vx, Vy        float64   // Velocity of the mouse
	Width, Height float64   // Dimensions of the mouse for collision
	Score         int       // Score of the player
	LastScoreTime time.Time // The last time a score was made
}

// NewMouse initializes a new mouse instance
func NewMouse() *Mouse {
	return &Mouse{
		X:             400, // Starting X position
		Y:             300, // Starting Y position
		Width:         10,  // Width of the mouse (not used when using image)
		Height:        10,  // Height of the mouse (not used when using image)
		Score:         0,   // Initial score
		LastScoreTime: time.Now().Add(-4 * time.Second), // Ensures there's no delay at start
	}
}

// ApplyPhysics applies simple physics to the velocity based on tilt and gravity
func ApplyPhysics(velocity, tilt, gravity float64) float64 {
	return velocity + tilt*gravity
}

// Update updates the mouse's position based on the tilt of the table
func (m *Mouse) Update(tiltX, tiltY float64, maze *Maze) {
	const gravity = 0.5

	// Check if we need to pause after scoring
	if time.Since(m.LastScoreTime) < 3*time.Second {
		return // Do not update position or velocity during the pause
	}

	// Apply physics to update velocity based on tilt
	m.Vx = ApplyPhysics(m.Vx, tiltX, gravity)
	m.Vy = ApplyPhysics(m.Vy, tiltY, gravity)

	// Update mouse position based on velocity
	m.X += m.Vx
	m.Y += m.Vy

	// Apply simple friction to slow down the mouse over time
	m.Vx *= 0.95
	m.Vy *= 0.95

	// Check for collisions with maze walls or holes
	if CheckCollision(m, maze) {
		// Collision with wall detected, bounce and decrement score
		m.Score = max(0, m.Score-1)
	}

	// Check for collisions with screen edges and bounce
	m.bounceOffEdges()
}

// bounceOffEdges checks if the mouse has hit the screen edges and makes it bounce
func (m *Mouse) bounceOffEdges() {
	// Left or right edge collision
	if m.X < 0 {
		m.X = 0
		m.Vx = -m.Vx // Reverse velocity to bounce
	} else if m.X+float64(graphics.MouseImage.Bounds().Dx()) > ScreenWidth {
		m.X = ScreenWidth - float64(graphics.MouseImage.Bounds().Dx())
		m.Vx = -m.Vx // Reverse velocity to bounce
	}

	// Top or bottom edge collision
	if m.Y < 0 {
		m.Y = 0
		m.Vy = -m.Vy // Reverse velocity to bounce
	} else if m.Y+float64(graphics.MouseImage.Bounds().Dy()) > ScreenHeight {
		m.Y = ScreenHeight - float64(graphics.MouseImage.Bounds().Dy())
		m.Vy = -m.Vy // Reverse velocity to bounce
	}
}

// Draw renders the mouse image on the screen and shows the score
func (m *Mouse) Draw(screen *ebiten.Image) {
	// Debug print to verify Draw function is called
	fmt.Printf("Drawing mouse at X: %.2f, Y: %.2f\n", m.X, m.Y)

	// Draw the mouse image
	if graphics.MouseImage != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(m.X, m.Y) // Translate to the mouse position
		screen.DrawImage(graphics.MouseImage, op)
	} else {
		// Fallback: Draw the mouse as a white rectangle if the image is missing
		ebitenutil.DrawRect(screen, m.X, m.Y, m.Width, m.Height, color.White)
	}

	// Draw the score on the right side of the screen using text.Draw
	scoreText := fmt.Sprintf("Score: %d", m.Score)
	textX := ScreenWidth - 150 // Position 150 pixels from the right edge
	textY := 30                // Position 30 pixels from the top
	text.Draw(screen, scoreText, fontFace, textX, textY, color.White)
}

// Maze represents the maze layout with walls and holes
type Maze struct {
	Walls []Wall // Slice of walls in the maze
	Holes []Hole // Slice of holes in the maze
}

// Wall represents a wall in the maze
type Wall struct {
	X, Y, Width, Height float64 // Position and size of the wall
}

// Hole represents a hole in the maze
type Hole struct {
	X, Y, Radius float64 // Position and size of the hole
}

// NewMaze initializes a new maze instance with predefined walls and holes
func NewMaze() *Maze {
	return &Maze{
		Walls: []Wall{
			{X: 50, Y: 50, Width: 100, Height: 20},
			{X: 200, Y: 150, Width: 20, Height: 100},
			// Add more walls as needed
		},
		Holes: []Hole{
			{X: 150, Y: 200, Radius: 15},
			{X: 300, Y: 100, Radius: 10},
			// Add more holes as needed
		},
	}
}

// Draw renders the maze on the screen
func (m *Maze) Draw(screen *ebiten.Image) {
	for _, wall := range m.Walls {
		// Draw each wall as a rectangle
		ebitenutil.DrawRect(screen, wall.X, wall.Y, wall.Width, wall.Height, color.Gray{Y: 128})
	}
	for _, hole := range m.Holes {
		// Draw each hole as a filled circle
		ebitenutil.DrawCircle(screen, hole.X, hole.Y, hole.Radius, color.Black)
	}
}

// CheckCollision checks if the mouse has collided with any walls or holes in the maze
func CheckCollision(mouse *Mouse, maze *Maze) bool {
	// Check collision with maze walls
	for _, wall := range maze.Walls {
		if mouse.X < wall.X+wall.Width && mouse.X+float64(graphics.MouseImage.Bounds().Dx()) > wall.X &&
			mouse.Y < wall.Y+wall.Height && mouse.Y+float64(graphics.MouseImage.Bounds().Dy()) > wall.Y {
			// Collision with a wall detected
			return true
		}
	}

	// Check if the mouse falls into any holes
	for _, hole := range maze.Holes {
		if math.Pow(mouse.X-hole.X, 2)+math.Pow(mouse.Y-hole.Y, 2) < math.Pow(hole.Radius, 2) {
			// Mouse has fallen into a hole
			mouse.Score++ // Increment the score by 1
			mouse.ResetPosition() // Reset the mouse position
			mouse.LastScoreTime = time.Now() // Record the time of scoring
			return false
		}
	}

	// No collision detected
	return false
}

// ResetPosition resets the mouse's position to the starting point
func (m *Mouse) ResetPosition() {
	m.X = 400
	m.Y = 300
	m.Vx = 0
	m.Vy = 0
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
