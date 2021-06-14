package container

import "log"

var DIC *Container

func init() {
	log.Println("container init")

	DIC = NewContainer()
}
