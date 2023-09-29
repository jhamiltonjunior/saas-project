package repositories

import (
    "my-saas-app/internal/domain/entities"
    "gorm.io/gorm"
)

type GormRepositoryImpl struct {
    db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepositoryImpl {
    return &GormRepositoryImpl{db: db}
}

func (g *GormRepositoryImpl) FindByID(id int) (*entities.User, error) {
    var user entities.User
    if err := g.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (g *GormRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
    var user entities.User
    if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (g *GormRepositoryImpl) Create(user *entities.User) error {
    return g.db.Create(user).Error
}

func (g *GormRepositoryImpl) Update(user *entities.User) error {
    return g.db.Save(user).Error
}

func (g *GormRepositoryImpl) Delete(id int) error {
    return g.db.Delete(&entities.User{}, id).Error
}