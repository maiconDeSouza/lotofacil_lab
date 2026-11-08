package model

import "time"

type Game struct {
	ID            string       `bson:"_id,omitempty"`
	ContestNumber string       `bson:"contest_number"`
	Strategy      StrategyType `bson:"strategy"`
	Numbers       Numbers      `bson:"numbers"`
	GeneratedAt   time.Time    `bson:"generated_at"`
	Hits          *int         `bson:"hits,omitempty"`
	Status        ResultStatus `bson:"status"`
}

type GameResult struct {
	GameID         string       `bson:"game_id"`
	ContestNumber  string       `bson:"contest_number"`
	Strategy       StrategyType `bson:"strategy"`
	Hits           int          `bson:"hits"`
	MatchedNumbers Numbers      `bson:"matched_numbers"`
	EvaluatedAt    time.Time    `bson:"evaluated_at"`
}
