package routers

import (
	"taufikRahadi/fiber-boilerplate/application/role/presenters"
	"taufikRahadi/fiber-boilerplate/common/ioc"

	"github.com/gofiber/fiber/v2"
)

type RoleV1Route struct {
	Module *ioc.Container
}

func NewRoleV1Route() ioc.BaseRouter {
	return RoleV1Route{
		Module: nil,
	}
}

func (r RoleV1Route) Init(module *ioc.Container) ioc.BaseRouter {
	r.Module = module

	return r
}

func (r RoleV1Route) InitRoute(app fiber.Router) {
	rolePresenter := presenters.NewRolePresenter(r.Module)

	app.Get("/v1/roles", rolePresenter.GetUserList)
}
