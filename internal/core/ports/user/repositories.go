package user

import (
	"AstroMap/internal/core/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	Get(user domain.User) error
}
