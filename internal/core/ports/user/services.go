package user

import "AstroMap/internal/core/domain"

type UserService interface {
	Registration(email string, pass string) (domain.RegistrationUser, error)
	Authorization(email string, pass string) (domain.AuthorizationUser, error)
}
