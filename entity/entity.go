package entity

import "github.com/google/uuid"

// ID type commonly used to uuid
type ID = uuid.UUID

// NewID Generates a new ID type
func NewID() ID {
	return ID(uuid.New())
}

// StringToID parse a string to ID type
func StringToID(idS string) (ID, error) {
	id, err := uuid.Parse(idS)
	return ID(id), err
}
