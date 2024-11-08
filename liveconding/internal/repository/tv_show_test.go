package repository

import (
	"layered_architecture/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test setup to initialize a mock DB connection for testing
func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.Episode{}, &domain.Location{}, &domain.Character{})
	return db
}

func TestCreateEpisodes(t *testing.T) {
	// Setup test DB
	db := setupTestDB()
	repo := NewTVShowRepository(db)

	// Test saving the episodes
	episodes := []domain.Episode{domain.Episode{
		OriginalID: 1,
		Name:       "Pilot",
		AirDate:    "December 2, 2013",
		Episode:    "S01E01",
		Characters: []string{"https://rickandmortyapi.com/api/character/1", "https://rickandmortyapi.com/api/character/2"},
		URL:        "https://rickandmortyapi.com/api/episode/1",
		Created:    "2017-11-10T12:56:33.798Z",
	}}
	episodesCreated, err := repo.CreateEpisodes(episodes)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 1, episodesCreated)
}

func TestCreateLocations(t *testing.T) {
	// Setup test DB
	db := setupTestDB()
	repo := NewTVShowRepository(db)

	// Test saving the locations
	locations := []domain.Location{domain.Location{
		OriginalID: 1,
		Name:       "Earth (C-137)",
		Type:       "Planet",
	}}
	locationsCreated, err := repo.CreateLocations(locations)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 1, locationsCreated)
}

func TestCreateCharacters(t *testing.T) {
	// Setup test DB
	db := setupTestDB()
	repo := NewTVShowRepository(db)

	// Test saving the characters
	characters := []domain.Character{domain.Character{
		OriginalID: 1,
		Name:       "Rick Sanchez",
		Gender:     "Male",
	}, domain.Character{
		OriginalID: 2,
		Name:       "Morty Smith",
		Gender:     "Male",
	}}
	charactersCreated, err := repo.CreateCharacters(characters)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 2, charactersCreated)
}
