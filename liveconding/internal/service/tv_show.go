package service

import (
	"fmt"
	"io"
	"layered_architecture/internal/domain"
	"layered_architecture/internal/repository"
	"layered_architecture/pkg/config"
	"log"
	"net/http"
	"sync"
)

type TVShowService interface {
	FetchFeeds() error
}

type TVShowFeedService struct {
	tvShowRepo repository.TVShowRepositoryDB
	config     *config.Config
}

func NewTVShowService(repo repository.TVShowRepositoryDB, cfg *config.Config) TVShowService {
	return &TVShowFeedService{tvShowRepo: repo, config: cfg}
}

func (s *TVShowFeedService) FetchFeeds() error {
	var wg sync.WaitGroup // Waits for a collection of goroutines
	wg.Add(3)

	go func() {
		defer wg.Done() // Decrements the counter
		feed := s.extractTVShows(fmt.Sprintf("%s/%s", s.config.RickAndMortyURL, s.config.EpisodePath))
		e := domain.NewEpisode()
		e.Deserialize(feed)
		episodesCreated, err := s.tvShowRepo.CreateEpisodes(e.Results)
		if err != nil {
			log.Printf("Error creating episodes on TVShowRepo: %+v", err)
		}
		log.Printf("Total number of episodes created: %d", episodesCreated)
	}()

	go func() {
		defer wg.Done()
		feed := s.extractTVShows(fmt.Sprintf("%s/%s", s.config.RickAndMortyURL, s.config.LocationPath))
		l := domain.NewLocation()
		l.Deserialize(feed)
		locationsCreated, err := s.tvShowRepo.CreateLocations(l.Results)
		if err != nil {
			log.Printf("Error creating locations on TVShowRepo: %+v", err)
		}
		log.Printf("Total number of locations created: %d", locationsCreated)
	}()

	go func() {
		defer wg.Done()
		feed := s.extractTVShows(fmt.Sprintf("%s/%s", s.config.RickAndMortyURL, s.config.CharacterPath))
		c := domain.NewCharacter()
		c.Deserialize(feed)
		charactersCreated, err := s.tvShowRepo.CreateCharacters(c.Results)
		if err != nil {
			log.Printf("Error creating character on TVShowRepo: %+v", err)
		}
		log.Printf("Total number of character created: %d", charactersCreated)
	}()

	wg.Wait() // Blocks until all goroutines finish
	return nil
}

func (s *TVShowFeedService) extractTVShows(url string) string {
	// FetchFeeds a new HTTP client
	client := &http.Client{}

	// FetchFeeds the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	return string(body)
}
