package infrastructure

import (
	"database/sql"
	"log"

	"github.com/shoooooman/sample-ddd-app/container"
	"github.com/shoooooman/sample-ddd-app/domain/model"
	"github.com/shoooooman/sample-ddd-app/domain/repository"
)

func init() {
	log.Println("infrastructure init")

	def := &container.Definition{
		Name: "UserRepository",
		Builder: func(c *container.Container) interface{} {
			db, ok := container.DIC.Inject("DB").(*sql.DB)
			if !ok {
				log.Fatal("injection type error")
			}
			return NewUserRepository(db)
		},
	}
	container.DIC.Register(def)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func NewTestUserRepository(db *sql.DB) repository.UserRepository {
	db, err := sql.Open("mysql", "root:password@tcp(host.docker.internal:3306)/test")
	if err != nil {
		panic(err)
	}
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) SelectByID(userID int) *model.User {
	row := u.db.QueryRow("select id, name from users where id = ?", userID)

	var (
		id   int
		name string
	)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		break
	default:
		log.Fatal(err)
	}

	user, err := model.NewUser(id, name)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (u *UserRepositoryImpl) SelectByName(userName string) *model.User {
	row := u.db.QueryRow("select id, name from users where name = ?", userName)

	var (
		id   int
		name string
	)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		break
	default:
		log.Fatal(err)
	}

	user, err := model.NewUser(id, name)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (u *UserRepositoryImpl) Insert(user *model.User) error {
	row := u.db.QueryRow("insert into users (id, name) values (?, ?) ", user.ID, user.Name)

	switch err := row.Scan(); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return nil
	default:
		return err
	}
}
