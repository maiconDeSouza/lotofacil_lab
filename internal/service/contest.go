package service

import (
	"net/http"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/lottery"
	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"github.com/maiconDeSouza/lotofacil_lab/internal/repository"
)

type ContestServiceInterface interface {
}

type ContestService struct {
	repoContest    repository.ContestRepositoryInterface
	repoAppSetting repository.AppSettingsRepositoryInterface
	env            *env.Env
}

func NewContestService(
	repoContest *repository.ContestRepositoryInterface,
	repoAppSetting *repository.AppSettingsRepositoryInterface,
	env *env.Env,
) *ContestService {
	return &ContestService{
		repoContest:    *repoContest,
		repoAppSetting: repoAppSetting,
		env:            env,
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

	a := model.AppSettings{}
}
