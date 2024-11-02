package repository

import (
	"errors"
	"log"

	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

type userRepo struct{
	db *gorm.DB
}

func NewUserRepo (db *gorm.DB) domain.UsersRepository{
  return &userRepo{
    db:db,
  }
}

func (m *userRepo) Create(user *domain.User) error{
  if err := m.db.Create(user).Error; err != nil{
    return err
  }
  return nil
}

func (m *userRepo) GetByEmail(email string) (*domain.User, error){
  user := &domain.User{}
  result := m.db.First(user, "email = ?", email)
  if result.Error != nil{
    if errors.Is(result.Error, gorm.ErrRecordNotFound){
      return user, nil
    }
    return user, result.Error
  }
  return user, nil
}

func (m *userRepo) Update(user *domain.User) error{
  log.Println(user)
  if err := m.db.Save(user).Error; err != nil{
    if(errors.Is(err, gorm.ErrRecordNotFound)){
      return err
    }
    return err
  }
  return nil
}