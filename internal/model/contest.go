package model

import "time"

type Prize struct {
	Description string  `bson:"description"`
	Tier        int     `bson:"tier"`
	Winners     int     `bson:"winners"`
	Value       float64 `bson:"value"`
}

type Contest struct {
	ID           string       `bson:"_id"`
	Number       int          `bson:"number"`
	Next         int          `bson:"next"`
	Lottery      string       `bson:"lottery"`
	DrawDate     time.Time    `bson:"draw_date"`
	Numbers      Numbers      `bson:"numbers"`
	Prize        []Prize      `bson:"prize"`
	Accumulated  bool         `bson:"accumulated"`
	ResultStatus ResultStatus `bson:"result_status"`
	CreatedAt    time.Time    `bson:"created_at"`
}

type Premio struct {
	Descricao  string  `json:"descricaoFaixa"`
	Faixa      int     `json:"faixa"`
	Ganhadores int     `json:"numeroDeGanhadores"`
	Valor      float64 `json:"valorPremio"`
}

type LotofacilRequestContest struct {
	Numero    int      `json:"numero"`
	Proximo   int      `json:"numeroConcursoProximo"`
	Data      string   `json:"dataApuracao"`
	Numeros   []string `json:"listaDezenas"`
	Premios   []Premio `json:"listaRateioPremio"`
	Acumulado bool     `json:"acumulado"`
}
