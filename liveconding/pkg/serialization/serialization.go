package serialization

import (
	"encoding/json"
	"log"
)

type Serializable interface {
	Serialize() string
	Deserialize(string) error
}

// Serialize is a generic function that serializes any type.
func Serialize[T any](v *T) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println("Error serializing:", err)
		return ""
	}
	return string(data)
}

// Deserialize is a generic function that deserializes a JSON string into any type.
func Deserialize[T any](data string, v *T) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		log.Println("Error deserializing:", err)
		return err
	}
	return nil
}
