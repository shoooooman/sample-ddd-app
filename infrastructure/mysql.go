package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shoooooman/sample-ddd-app/container"
)

func init() {
	def := &container.Definition{
		Name: "DB",
		Builder: func(c *container.Container) interface{} {
			return newDB()
		},
	}
	container.DIC.Register(def)
}

func newDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(host.docker.internal:3306)/sample")
	if err != nil {
		panic(err)
	}
	return db
}
