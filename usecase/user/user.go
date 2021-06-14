package usecase

import (
	"fmt"
	"log"

	"github.com/shoooooman/sample-ddd-app/container"
	"github.com/shoooooman/sample-ddd-app/domain/model"
	"github.com/shoooooman/sample-ddd-app/domain/repository"
	"github.com/shoooooman/sample-ddd-app/domain/service"
	_ "github.com/shoooooman/sample-ddd-app/infrastructure" // to register UserRepository into the container
)

func init() {
	log.Println("usecase init")

	def := &container.Definition{
		Name: "UserUsecase",
		Builder: func(c *container.Container) interface{} {
			userRepository, _ := c.Inject("UserRepository").(repository.UserRepository)
			return NewUserUsecase(userRepository)
		},
	}
	container.DIC.Register(def)
}

type UserUsecase interface {
	FindUser(int) *model.User
	CreateUser(int, string) (int, error)
}

type UserUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepository: repository}
}

func (u *UserUsecaseImpl) FindUser(userID int) *model.User {
	return u.userRepository.SelectByID(userID)
}

func (u *UserUsecaseImpl) CreateUser(userID int, name string) (int, error) {
	user, err := model.NewUser(userID, name)
	if err != nil {
		return 0, err
	}

	userService, _ := container.DIC.Inject("UserService").(service.UserService)
	if userService.Exists(user) {
		return 0, fmt.Errorf("Same user name exists")
	}

	return u.userRepository.Insert(user), nil
}
