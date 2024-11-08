package domain

import (
	"layered_architecture/pkg/serialization"
)

type Episode struct {
	OriginalID      int      `json:"id"`
	Name            string   `json:"name"`
	AirDate         string   `json:"air_date"`
	Episode         string   `json:"episode"`
	Characters      []string `json:"characters" gorm:"-"`
	CharactersPlain string   `json:"Characters_plain"`
	URL             string   `json:"url"`
	Created         string   `json:"created"`
}

type EpisodeResponse struct {
	Info              // Composition
	Results []Episode `json:"results, omitempty"`
}

// Constructor for creating a new instance (idiomatic approach)
func NewEpisode() *EpisodeResponse {
	return &EpisodeResponse{}
}

// Serialize method converts the struct to a JSON string representation
func (u *EpisodeResponse) Serialize() string {
	return serialization.Serialize[EpisodeResponse](u)
}

// Deserialize method converts a JSON string back into the struct
func (u *EpisodeResponse) Deserialize(data string) error {
	return serialization.Deserialize[EpisodeResponse](data, u)
}

// Ensures that Struct implements the interface at compile time
var _ serialization.Serializable = (*EpisodeResponse)(nil)
