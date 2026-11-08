package model

import "time"

type Contest struct {
	ID           string       `bson:"_id"`
	Next         string       `bson:"number"`
	Lottery      string       `bson:"lottery"`
	DrawDate     time.Time    `bson:"draw_date"`
	Numbers      Numbers      `bson:"numbers"`
	Prize        float64      `bson:"prize"`
	Accumulated  bool         `bson:"accumulated"`
	ResultStatus ResultStatus `bson:"result_status"`
	CreatedAt    time.Time    `bson:"created_at"`
	UpdatedAt    time.Time    `bson:"updated_at"`
}
