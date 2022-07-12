package user

import (
	"AstroMap/internal/core/domain"
	"AstroMap/internal/core/ports/user"
)

type service struct {
	userRepository user.UserRepository
}

func New(userRepository user.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func Registration(email string, pass string) domain.RegistrationUser {
	return domain.RegistrationUser{
		Email:     email,
		Passsword: pass,
	}
}
