package config

// Example config structure for game settings
type Config struct {
	ScreenWidth  int
	ScreenHeight int
	Difficulty   string
}

// LoadConfig loads game configuration settings
func LoadConfig() *Config {
	// Load settings from a file or environment variables
	return &Config{
		ScreenWidth:  800,
		ScreenHeight: 600,
		Difficulty:   "normal",
	}
}
