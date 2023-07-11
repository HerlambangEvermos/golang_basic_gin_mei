package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Usename  string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(reqPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPass))

	if err != nil {
		return err
	}

	return nil
}
