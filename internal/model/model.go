package model

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
}

type Game struct {
	id          uuid.UUID
	giantbombId string
}
