package service

import (
	"log"

	"github.com/shoooooman/sample-ddd-app/container"
	"github.com/shoooooman/sample-ddd-app/domain/model"
	"github.com/shoooooman/sample-ddd-app/domain/repository"
	_ "github.com/shoooooman/sample-ddd-app/infrastructure" // to register UserRepository into the container
)

func init() {
	log.Println("service init")

	def := &container.Definition{
		Name: "UserService",
		Builder: func(c *container.Container) interface{} {
			userRepository, ok := c.Inject("UserRepository").(repository.UserRepository)
			if !ok {
				log.Fatal("injection type error")
			}
			return NewUserService(userRepository)
		},
	}
	container.DIC.Register(def)
}

type UserService interface {
	Exists(*model.User) bool
}

type UserServiceImpl struct {
	repository repository.UserRepository
}

func (u *UserServiceImpl) Exists(user *model.User) bool {
	return u.repository.SelectByName(user.Name) != nil
}

func NewUserService(repository repository.UserRepository) UserService {
	return &UserServiceImpl{repository}
}
