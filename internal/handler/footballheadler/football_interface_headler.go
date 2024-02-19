package footballheadler

import (
	"github.com/odanaraujo/futebol-api/internal/handler/footballservice"
	"net/http"
)

func NewFootballHeadler(service footballservice.FootballService) FootballHeadler {
	return &handler{service: service}
}

type handler struct {
	service footballservice.FootballService
}

type FootballHeadler interface {
	GetTablePernambucanoA1(w http.ResponseWriter, r *http.Request)
}
