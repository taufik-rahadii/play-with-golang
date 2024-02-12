package ioc

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type BaseObject interface {
	Init(module *Container) BaseObject
}

type BaseRouter interface {
	Init(module *Container) BaseRouter
	InitRoute(app fiber.Router)
}

type BasePresenter interface {
	Init(module *Container) BasePresenter
}

type Container struct {
	Deps       map[string]interface{}
	Providers  map[string]BaseObject
	Presenters map[string]BaseRouter
	Exports    map[string]func() BaseObject
}

func (c *Container) injectProviders(providers map[string]func() BaseObject) {
	for key, val := range providers {
		c.Providers[key] = val().Init(c)
	}
}

func (c *Container) CallProviders(name string, structed BaseObject) (BaseObject, error) {
	providers := c.Providers[name]

	if providers == nil {
		return nil, errors.New("there is no service provider registered with this name: " + name)
	}

	return providers, nil
}

func (c *Container) injectPresenters(presenters map[string]func() BaseRouter) {
	for key, val := range presenters {
		c.Presenters[key] = val().Init(c)
	}
}

func (c *Container) exposePresenters(app fiber.Router) {
	for _, val := range c.Presenters {
		val.InitRoute(app)
	}
}

func (c *Container) exportProviders() {

}

func NewContainer(app fiber.Router, imports map[string]interface{}, providers map[string]func() BaseObject, presenters map[string]func() BaseRouter, exports map[string]func() BaseObject) *Container {
	c := &Container{
		Deps:       imports,
		Providers:  map[string]BaseObject{},
		Presenters: map[string]BaseRouter{},
		Exports:    exports,
	}

	c.injectProviders(providers)
	c.injectPresenters(presenters)
	c.exposePresenters(app)

	return c
}
