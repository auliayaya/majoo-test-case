package domain

import (
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint      `gorm:"primaryKey;column:id;"`
	Name      string    `gorm:"column:name"`
	Username  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	Merchants Merchant  `gorm:"foreignkey:UserID;references:ID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy uint      `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy uint      `gorm:"column:updated_by"`
}

func (User) TableName() string {
	return "Users"
}

type (
	UserRepository interface {
		GetUserByUsername(username string) (*User, error)
		CreateUser(*User) (*User, error)
	}
	UserUsecase interface {
		UserLogin(dto.LoginRequestDTO) (models.Responses, error)
		UserRegister(dto.RegisterRequestDTO) (models.Responses, error)
	}
	UserHandler interface {
		Login(*gin.Context)
		Register(*gin.Context)
	}
)
