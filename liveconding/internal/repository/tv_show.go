package repository

import (
	"layered_architecture/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type TVShowRepositoryDB interface {
	CreateEpisodes(episodes []domain.Episode) (int, error)
	CreateLocations(locations []domain.Location) (int, error)
	CreateCharacters(characters []domain.Character) (int, error)
}

type TVShowRepository struct {
	db *gorm.DB
}

var _ TVShowRepositoryDB = (*TVShowRepository)(nil)

func NewTVShowRepository(db *gorm.DB) TVShowRepositoryDB {
	return &TVShowRepository{db: db}
}

func (r *TVShowRepository) CreateEpisodes(episodes []domain.Episode) (int, error) {
	count := 0
	for _, episode := range episodes {
		episode.CharactersPlain = strings.Join(episode.Characters, ",")
		if err := r.db.Create(episode).Error; err != nil {
			return count, err // Return count and error if creation fails
		}
		count++ // Increment the counter if creation is successful

	}
	return count, nil
}

func (r *TVShowRepository) CreateLocations(locations []domain.Location) (int, error) {
	count := 0
	for _, location := range locations {
		location.ResidentsPlain = strings.Join(location.Residents, ",")
		if err := r.db.Create(location).Error; err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (r *TVShowRepository) CreateCharacters(characters []domain.Character) (int, error) {
	count := 0
	for _, character := range characters {
		character.EpisodePlain = strings.Join(character.Episode, ",")
		if err := r.db.Create(character).Error; err != nil {
			return count, err
		}
		count++
	}
	return count, nil

}
