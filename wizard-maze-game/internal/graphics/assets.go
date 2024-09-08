package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

var (
	MouseImage      *ebiten.Image
	WallImage       *ebiten.Image
	HoleImage       *ebiten.Image
	BackgroundImage *ebiten.Image
)

func init() {
	var err error

	// Load mouse image
	MouseImage, _, err = ebitenutil.NewImageFromFile("assets/images/mouse.png")
	if err != nil {
		log.Fatalf("Failed to load mouse image: %v", err)
	}

	// Load wall image
	WallImage, _, err = ebitenutil.NewImageFromFile("assets/images/wall.png")
	if err != nil {
		log.Fatalf("Failed to load wall image: %v", err)
	}

	// Load hole image
	HoleImage, _, err = ebitenutil.NewImageFromFile("assets/images/hole.png")
	if err != nil {
		log.Fatalf("Failed to load hole image: %v", err)
	}

	// Load background image
	BackgroundImage, _, err = ebitenutil.NewImageFromFile("assets/images/background.png")
	if err != nil {
		log.Fatalf("Failed to load background image: %v", err)
	}
}

// RenderBackground renders the background image or fills the screen with a color
func RenderBackground(screen *ebiten.Image, width, height int) {
	if BackgroundImage != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(width)/float64(BackgroundImage.Bounds().Dx()), float64(height)/float64(BackgroundImage.Bounds().Dy()))
		screen.DrawImage(BackgroundImage, op)
	} else {
		// If no background image, fill with a solid color
		screen.Fill(color.RGBA{0x11, 0x11, 0x33, 0xff}) // Example solid color fill
	}
}

// DrawMouse draws the mouse image on the screen
func DrawMouse(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(MouseImage, op)
}

// DrawWall draws the wall image on the screen
func DrawWall(screen *ebiten.Image, x, y, width, height float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(width/float64(WallImage.Bounds().Dx()), height/float64(WallImage.Bounds().Dy()))
	op.GeoM.Translate(x, y)
	screen.DrawImage(WallImage, op)
}

// DrawHole draws the hole image on the screen
func DrawHole(screen *ebiten.Image, x, y, radius float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2*radius/float64(HoleImage.Bounds().Dx()), 2*radius/float64(HoleImage.Bounds().Dy()))
	op.GeoM.Translate(x-radius, y-radius)
	screen.DrawImage(HoleImage, op)
}
