package usecase

import (
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
			userRepository, ok := c.Inject("UserRepository").(repository.UserRepository)
			if !ok {
				log.Fatal("injection type error")
			}
			return NewUserUsecase(userRepository)
		},
	}
	container.DIC.Register(def)
}

type UserUsecase interface {
	FindUser(int) *model.User
	CreateUser(int, string) error
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

func (u *UserUsecaseImpl) CreateUser(userID int, name string) error {
	user, err := model.NewUser(userID, name)
	if err != nil {
		return &ModelError{err.Error()}
	}

	userService, ok := container.DIC.Inject("UserService").(service.UserService)
	if !ok {
		log.Fatal("injection type error")
	}
	if userService.Exists(user) {
		return &ModelError{"same user name exists"}
	}

	if err := u.userRepository.Insert(user); err != nil {
		return &RepositoryError{err.Error()}
	}

	return nil
}
