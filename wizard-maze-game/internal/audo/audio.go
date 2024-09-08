package audio

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
	"os"
)

var (
	CollisionSound *audio.Player
	ScoreSound     *audio.Player
	GameOverSound  *audio.Player
	AudioContext   *audio.Context
)

func init() {
	// Initialize audio context
	AudioContext = audio.NewContext(44100)

	// Load sounds
	CollisionSound = loadSound("assets/sounds/collision-sound.wav")
	ScoreSound = loadSound("assets/sounds/score-sound.wav")
	GameOverSound = loadSound("assets/sounds/game-over.wav")
}

func loadSound(path string) *audio.Player {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open sound file: %s, %v", path, err)
	}
	defer f.Close()

	d, err := wav.Decode(AudioContext, f)
	if err != nil {
		log.Fatalf("Failed to decode sound file: %s, %v", path, err)
	}

	p, err := audio.NewPlayer(AudioContext, d)
	if err != nil {
		log.Fatalf("Failed to create sound player: %s, %v", path, err)
	}

	return p
}

func PlayCollisionSound() {
	CollisionSound.Rewind()
	CollisionSound.Play()
}

func PlayScoreSound() {
	ScoreSound.Rewind()
	ScoreSound.Play()
}

func PlayGameOverSound() {
	GameOverSound.Rewind()
	GameOverSound.Play()
}
