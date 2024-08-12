package repositories

import (
	"my-saas-app/src/domain/entities"
)

type Period struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

type OnlyYearPeriod struct {
	Year int `json:"year"`
}

type RemunerationRepository interface {
	FindByID(int int) (*entities.Remuneration, error)
	FindAll(*Period) ([]entities.Remuneration, error)
	FindAllByYear(*OnlyYearPeriod) ([]entities.Remuneration, error)
	FindByName(cnpj string) (*entities.Remuneration, error)
	FindByType(cnpj string) (*entities.Remuneration, error)
	Create(remuneration *entities.Remuneration) (int, error)
	Update(remuneration *entities.Remuneration) error
	Delete(remuneration *entities.Remuneration) error
}
