package service

import (
	"net/http"
	"slices"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/lottery"
	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"github.com/maiconDeSouza/lotofacil_lab/internal/repository"
	"github.com/maiconDeSouza/lotofacil_lab/pkg/random"
)

type ContestServiceInterface interface {
	Start() (*model.GamesNow, error)
}

type ContestService struct {
	repoContest         repository.ContestRepositoryInterface
	repoAppSetting      repository.AppSettingsRepositoryInterface
	repoGame            repository.GameRepositoryInterface
	repoNumberFrequency repository.NumberFrequencyInterface
	env                 *env.Env
}

func NewContestService(
	repoContest repository.ContestRepositoryInterface,
	repoAppSetting repository.AppSettingsRepositoryInterface,
	repoGame repository.GameRepositoryInterface,
	repoNumberFrequency repository.NumberFrequencyInterface,
	env *env.Env,
) *ContestService {
	return &ContestService{
		repoContest:         repoContest,
		repoAppSetting:      repoAppSetting,
		repoGame:            repoGame,
		repoNumberFrequency: repoNumberFrequency,
		env:                 env,
	}
}

func (s *ContestService) Start() (*model.GamesNow, error) {
	lastContest := model.Contest{}
	for {
		if err := s.repoContest.LastContest(&lastContest); err != nil {
			return nil, err
		}

		c, status := lottery.GetContest(lastContest.Next, s.env)
		if status == http.StatusOK {
			if err := s.repoContest.CreateContest(c); err != nil {
				return nil, err
			}
			continue
		}
		break
	}

	app := model.AppSettings{}
	if err := s.repoAppSetting.FindAppSettings(&app); err != nil {
		return nil, err
	}

	f := model.Game{
		ContestNumber: lastContest.Next,
		Strategy:      model.StrategyFixed,
		Numbers:       app.FixedGame,
		GeneratedAt:   time.Now(),
		Status:        model.ResultWaiting,
	}

	if err := s.repoGame.CreateGame(&f); err != nil {
		return nil, err
	}

	var contests []*model.Contest
	if err := s.repoContest.FindAllContest(&contests); err != nil {
		return nil, err
	}

	slices.SortFunc(contests, func(a, b *model.Contest) int {
		if a.Number > b.Number {
			return -1
		}
		if a.Number < b.Number {
			return 1
		}
		return 0
	})

	lastContests := contests[:30]

	var numbers []*model.NumberFrequency

	if err := s.repoNumberFrequency.FindAllNumberFrequency(&numbers); err != nil {
		return nil, err
	}

	for _, v := range numbers {
		v.Count = 0
	}

	for _, v := range lastContests {
		for _, n := range v.Numbers {
			for _, i := range numbers {
				if i.Number == n {
					i.Count++
					break
				}
			}
		}
	}

	for _, v := range numbers {
		if err := s.repoNumberFrequency.UpdateNumberFrequency(v); err != nil {
			return nil, err
		}
	}

	slices.SortFunc(numbers, func(a, b *model.NumberFrequency) int {
		if a.Count > b.Count {
			return -1
		}
		if a.Count < b.Count {
			return 1
		}
		return 0
	})

	frequents, err := random.GenerateFrequent(&numbers)
	if err != nil {
		return nil, err
	}

	nf := model.Game{
		ContestNumber: lastContest.Next,
		Strategy:      model.StrategyFrequency,
		Numbers:       frequents,
		GeneratedAt:   time.Now(),
		Status:        model.ResultWaiting,
	}

	nr := model.Game{
		ContestNumber: lastContest.Next,
		Strategy:      model.StrategyRandom,
		Numbers:       random.GenerateLotofacilNumbers15(),
		GeneratedAt:   time.Now(),
		Status:        model.ResultWaiting,
	}

	gamersNow := model.GamesNow{
		Fixed:           f,
		Frequency:       nf,
		Random:          nr,
		NumberFrequency: numbers,
	}

	return &gamersNow, nil
}
