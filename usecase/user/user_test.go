package usecase

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/shoooooman/sample-ddd-app/container"
	"github.com/shoooooman/sample-ddd-app/domain/model"
	_ "github.com/shoooooman/sample-ddd-app/infrastructure" // to register UserRepository into the container
)

func init() {
	log.Println("usecase test init")

	def := &container.Definition{
		Name: "DB",
		Builder: func(c *container.Container) interface{} {
			return newTestDB()
		},
	}
	container.DIC.Register(def)
}

func newTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(host.docker.internal:3306)/test")
	if err != nil {
		panic(err)
	}
	return db
}

func TestMain(m *testing.M) {
	// TODO: make a transaction

	code := m.Run()

	// TODO: rollback the transaction

	os.Exit(code)
}

func TestFindUser(t *testing.T) {
	userUsecase, _ := container.DIC.Inject("UserUsecase").(UserUsecase)

	var tests = []struct {
		name     string
		expected *model.User
		userID   int
	}{
		{"Userが存在する場合", &model.User{ID: 1, Name: "shoma"}, 1},
		{"Userが存在しない場合", nil, -1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := userUsecase.FindUser(tt.userID)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("userID(%v): expected %v, actual %v", tt.userID, tt.expected, actual)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	userUsecase, _ := container.DIC.Inject("UserUsecase").(UserUsecase)

	var tests = []struct {
		name     string
		expected int
		userID   int
		userName string
	}{
		{"Userを作成できる場合", 1, 1, "shoma"},
		{"Userが既に作成されている場合", 0, 1, "shoma"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := userUsecase.CreateUser(tt.userID, tt.userName)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("userID(%v): expected %v, actual %v", tt.userID, tt.expected, actual)
			}
		})
	}
}
