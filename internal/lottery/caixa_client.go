package lottery

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
)

var url = "https://servicebus2.caixa.gov.br/portaldeloterias/api/lotofacil"

func GetLastContest() (int, int) {
	type Lotofacil struct {
		Numero int `json:"numero"`
	}
	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição: ", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao acessar a API: ", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, res.StatusCode
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Erro ao ler resposta da API: ", err)
	}

	l := Lotofacil{}

	if err := json.Unmarshal(body, &l); err != nil {
		log.Fatal("API retornou json inesperado: ", err)
	}

	return l.Numero, http.StatusOK
}

func GetContest(number int, env *env.Env) (*model.Contest, int) {
	type Lotofacil struct {
		Numero int `json:"numero"`
	}
	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", url, number), nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição: ", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao acessar a API: ", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Erro ao ler resposta da API: ", err)
	}

	l := model.LotofacilRequestContest{}

	if err := json.Unmarshal(body, &l); err != nil {
		log.Fatal("API retornou json inesperado: ", err)
	}

	layout := "02/01/2006"
	d, err := time.Parse(layout, l.Data)
	if err != nil {
		log.Fatal("Erro ao converter data: ", err)
	}

	var p []model.Prize

	for _, v := range l.Premios {
		n := model.Prize{
			Description: v.Descricao,
			Tier:        v.Faixa,
			Winners:     v.Ganhadores,
			Value:       v.Valor,
		}

		p = append(p, n)
	}

	var numbers model.Numbers

	for _, v := range l.Numeros {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Erro ao converter string dos resultaos em números: ", err)
		}

		numbers = append(numbers, n)
	}

	c := &model.Contest{
		ID:           fmt.Sprintf("%d", l.Numero),
		Number:       l.Numero,
		Next:         l.Proximo,
		Lottery:      env.Lottery.Name,
		DrawDate:     d,
		Numbers:      numbers,
		Prize:        p,
		Accumulated:  l.Acumulado,
		ResultStatus: model.ResultConfirmed,
		CreatedAt:    time.Now(),
	}

	return c, http.StatusOK
}
