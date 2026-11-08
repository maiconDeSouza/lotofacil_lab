package model

import "time"

type AppSettings struct {
	ID        string    `bson:"_id"`
	FixedGame Numbers   `bson:"fixed_game"`
	CreatedAt time.Time `bson:"created_at"`
}
