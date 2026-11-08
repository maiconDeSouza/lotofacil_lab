package model

type Numbers []int

type StrategyType string

const (
	StrategyFixed     StrategyType = "fixed"
	StrategyFrequency StrategyType = "frequency"
	StrategyRandom    StrategyType = "random"
)

type ResultStatus string

const (
	ResultConfirmed ResultStatus = "confirmed"
	ResultWaiting   ResultStatus = "waiting"
)

type NumberGroup struct {
	Position int     `json:"position"`
	Numbers  Numbers `json:"numbers"`
}
