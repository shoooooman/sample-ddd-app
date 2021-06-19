package model

import "fmt"

type User struct {
	ID   int
	Name string
}

const minNameLen = 5

func NewUser(id int, name string) (*User, error) {
	if len(name) < minNameLen {
		return nil, fmt.Errorf("name is too short")
	}

	return &User{
		ID:   id,
		Name: name,
	}, nil
}
