package footballservice

import "github.com/odanaraujo/futebol-api/internal/handler/footballrepository"

func NewFootballService(repo footballrepository.FootballRepository) FootballService {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo footballrepository.FootballRepository
}

type FootballService interface {
	GetTablePernambucanoA1() error
}
