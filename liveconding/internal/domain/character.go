package domain

import "layered_architecture/pkg/serialization"

type Character struct {
	OriginalID   int      `json:"id"`
	Name         string   `json:"name"`
	Status       string   `json:"status"`
	Species      string   `json:"species"`
	Type         string   `json:"type"`
	Gender       string   `json:"gender"`
	Image        string   `json:"image"`
	Episode      []string `json:"episode" gorm:"-"`
	EpisodePlain string   `json:"episode_plain"`
	URL          string   `json:"url"`
	Created      string   `json:"created"`
}

type CharacterLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CharacterResponse struct {
	Info                // Composition
	Results []Character `json:"results"`
}

// Constructor for creating a new instance (idiomatic approach)
func NewCharacter() *CharacterResponse {
	return new(CharacterResponse)
}

// Serialize method converts the struct to a JSON string representation
func (u *CharacterResponse) Serialize() string {
	return serialization.Serialize[CharacterResponse](u)
}

// Deserialize method converts a JSON string back into the struct
func (u *CharacterResponse) Deserialize(data string) error {
	return serialization.Deserialize[CharacterResponse](data, u)
}

// Ensures that Struct implements the interface at compile time
var _ serialization.Serializable = (*CharacterResponse)(nil)
