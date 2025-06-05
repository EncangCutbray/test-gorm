package request

import "cutbray/test-gorm/models"

type CreateUserRequest struct {
	Username string `validate:"required,min=5"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

func (r *CreateUserRequest) ToUserModel() *models.UserModel {
	return &models.UserModel{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
