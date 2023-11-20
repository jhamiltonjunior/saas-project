package database

import (
	"my-saas-app/internal/domain/entities"

	"gorm.io/gorm"
)

type GormUserRepositoryImpl struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepositoryImpl {
	return &GormUserRepositoryImpl{db: db}
}

func (g *GormUserRepositoryImpl) FindByID(id int) (*entities.User, error) {
	var user entities.User
	if err := g.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormUserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormUserRepositoryImpl) Create(user *entities.User) (int, error) {
	userDB := g.db.Create(user)
	return user.ID, userDB.Error
}

func (g *GormUserRepositoryImpl) Update(user *entities.User) error {
	return g.db.Save(user).Error
}

func (g *GormUserRepositoryImpl) Delete(id int) error {
	return g.db.Delete(&entities.User{}, id).Error
}
