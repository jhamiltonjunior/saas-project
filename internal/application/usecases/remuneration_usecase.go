package usecases

import (
	"errors"
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"
)

type RemunerationUseCase struct {
	remunerationRepository repositories.RemunerationRepository
}

func NewRemunerationUseCase(remunerationRepository repositories.RemunerationRepository) *RemunerationUseCase {
	return &RemunerationUseCase{
		remunerationRepository: remunerationRepository,
	}
}

func (bk *RemunerationUseCase) Create(input *entities.Remuneration) (int, error) {
	remuneration := &entities.Remuneration{
		Name:         input.Name,
		Value:        input.Value,
		UserID:       input.UserID,
		RecurrenceID: input.RecurrenceID,
	}

	remunerationId, err := bk.remunerationRepository.Create(remuneration)

	if err != nil {
		return 0, err
	}

	return remunerationId, nil
}

func (bk *RemunerationUseCase) GetRemunerationByID(id int) (*entities.Remuneration, error) {
	return bk.remunerationRepository.FindByID(id)
}

func (bk *RemunerationUseCase) GetRemunerationByName(name string) (*entities.Remuneration, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return bk.remunerationRepository.FindByName(name)
}

func (bk *RemunerationUseCase) GetAllRemunerationByMonths(period *repositories.Period) ([]entities.Remuneration, error) {
	return bk.remunerationRepository.FindAll(period)
}

func (bk *RemunerationUseCase) GetAllRemunerationByYear(year *repositories.OnlyYearPeriod) ([]entities.Remuneration, error) {
	return bk.remunerationRepository.FindAllByYear(year)
}

func (bk *RemunerationUseCase) Update(input *entities.Remuneration) error {
	remuneration := &entities.Remuneration{
		ID:     input.ID,
		Name:   input.Name,
		Value:  input.Value,
		UserID: input.UserID,
	}

	return bk.remunerationRepository.Update(remuneration)
}

func (bk *RemunerationUseCase) Delete(remuneration *entities.Remuneration) error {
	return bk.remunerationRepository.Delete(remuneration)
}
