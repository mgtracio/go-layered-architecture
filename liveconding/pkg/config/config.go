package config

import (
	"layered_architecture/internal/domain"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	RickAndMortyURL string `envconfig:"RICK_AND_MORTY_URL" default:"https://rickandmortyapi.com/api"`
	CharacterPath   string `envconfig:"CHARACTERS_PATH" default:"character"`
	LocationPath    string `envconfig:"LOCATIONS_PATH" default:"location"`
	EpisodePath     string `envconfig:"EPISODES_PATH" default:"episode"`
	AppPort         int    `envconfig:"APP_PORT" default:"8080"`
}

// Singleton instance and sync to ensure thread-safety
var (
	instance *Config
	once     sync.Once
)

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		envconfig.MustProcess("", instance)
	})
	return instance
}

// GetConfig returns the singleton instance of the Config.
func GetConfig() *Config {
	return LoadConfig()
}

func InitDB() *gorm.DB {
	wd, _ := os.Getwd() // Get the current working directory
	dbPath := filepath.Join(wd, "../../db/test.ldb")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&domain.Episode{}, &domain.Location{}, &domain.Character{})
	if err != nil {
		log.Fatalf("Error migrating databas: %v", err)
	}

	return db
}
