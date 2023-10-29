package service

import (
	"time"

	"github.com/fnunezzz/go-project-template/internal/repository"
)


type appService struct {
	dao repository.DAO
}

type AppService interface {
	HealthCheck() (time.Time, error)
}

func NewAppService(dao repository.DAO) AppService {
	return &appService{dao: dao}
}

func (u *appService) HealthCheck() (time.Time, error) {
	return u.dao.NewAppRepository().GetCurrentTimestamp()
}