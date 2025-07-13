package model

import "time"

type Token struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Exchange  string
}

type AppMode struct {
	Mode bool
}
