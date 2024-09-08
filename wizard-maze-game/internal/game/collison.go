package game

import (
	"wizard-maze-game/internal/entities"
)

// CheckCollision checks if the mouse has collided with any walls or holes in the maze.
func CheckCollision(mouse *entities.Mouse, maze *entities.Maze) bool {
	// Check collision with maze walls
	for _, wall := range maze.Walls {
		if mouse.X < wall.X+wall.Width && mouse.X+mouse.Width > wall.X &&
			mouse.Y < wall.Y+wall.Height && mouse.Y+mouse.Height > wall.Y {
			// Collision with a wall detected
			return true
		}
	}

	// Check if the mouse falls into any holes
	for _, hole := range maze.Holes {
		if (mouse.X-hole.X)*(mouse.X-hole.X)+(mouse.Y-hole.Y)*(mouse.Y-hole.Y) < hole.Radius*hole.Radius {
			// Mouse has fallen into a hole
			return true
		}
	}

	// No collision detected
	return false
}
