package model

import "time"

type NumberFrequency struct {
	ID     string `bson:"_id"`
	Number int    `bson:"number"`
	Count  int    `bson:"count"`
}

type FrequencyAnalysis struct {
	ID             string            `bson:"_id,omitempty"`
	TargetContest  string            `bson:"target_contest"`
	SourceContests []int             `bson:"source_contests"`
	Frequencies    []NumberFrequency `bson:"frequencies"`
	Groups         []NumberGroup     `bson:"groups"`
	CreatedAt      time.Time         `bson:"created_at"`
}
