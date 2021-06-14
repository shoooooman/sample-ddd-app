package container

import "log"

type Builder func(*Container) interface{}

type Definition struct {
	Name    string
	Builder Builder
}

type Container struct {
	store map[string]Builder
}

func NewContainer() *Container {
	return &Container{
		store: make(map[string]Builder),
	}
}

func (c *Container) Register(d *Definition) {
	c.store[d.Name] = d.Builder
}

func (c *Container) Inject(key string) interface{} {
	builder, ok := c.store[key]
	if !ok {
		log.Fatalf("%v is not registered", key)
	}
	instance := builder(c)
	return instance
}
