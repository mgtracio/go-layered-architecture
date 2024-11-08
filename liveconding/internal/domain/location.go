package domain

import "layered_architecture/pkg/serialization"

type Location struct {
	OriginalID     int      `json:"id"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Dimension      string   `json:"dimension"`
	Residents      []string `json:"residents" gorm:"-"`
	ResidentsPlain string   `json:"residents_plain"`
	URL            string   `json:"url"`
	Created        string   `json:"created"`
}

type LocationResponse struct {
	Info               // Composition
	Results []Location `json:"results,omitempty"`
}

// Constructor for creating a new instance (idiomatic approach)
func NewLocation() *LocationResponse {
	return &LocationResponse{}
}

// Serialize method converts the struct to a JSON string representation
func (u *LocationResponse) Serialize() string {
	return serialization.Serialize[LocationResponse](u)
}

// Deserialize method converts a JSON string back into the struct
func (u *LocationResponse) Deserialize(data string) error {
	return serialization.Deserialize[LocationResponse](data, u)
}

// Ensures that Struct implements the interface at compile time
var _ serialization.Serializable = (*LocationResponse)(nil)
