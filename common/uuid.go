package common

import "github.com/google/uuid"

func GenUUID() uuid.UUID {
	new_id, _ := uuid.NewV7()
	return new_id
}

func ParseUUID(s string) uuid.UUID {
	id := uuid.MustParse(s)
	return id
}
