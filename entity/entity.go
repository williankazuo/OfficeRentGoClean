package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func StringToID(idS string) (ID, error) {
	id, err := uuid.Parse(idS)
	return ID(id), err
}
