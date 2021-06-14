package repository

import "github.com/shoooooman/sample-ddd-app/domain/model"

type UserRepository interface {
	SelectByID(int) *model.User
	SelectByName(string) *model.User
	Insert(*model.User) error
}
