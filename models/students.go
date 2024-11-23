package models

import "time"

type Student struct {
	Name         string
	Email        string
	Age          int
	RegisteredAt time.Time
}
