package port

import "github.com/xXHachimanXx/go-microservice-starter/internal/core/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindAll() ([]domain.User, error)
	FindByID(id int) (domain.User, error)
	Delete(id int) error
}
