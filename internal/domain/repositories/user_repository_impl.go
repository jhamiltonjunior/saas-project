package repositories

import (
    "my-saas-app/internal/domain/entities"
    "gorm.io/gorm"
)

type UserRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindByID(id int) (*entities.User, error) {
    var user entities.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
    var user entities.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
    return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
    return r.db.Delete(&entities.User{}, id).Error
}