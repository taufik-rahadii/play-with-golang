package application

import (
	"taufikRahadi/fiber-boilerplate/application/role"
	"taufikRahadi/fiber-boilerplate/common/ioc"

	"github.com/gofiber/fiber/v2"
)

func ApplicationModule(app fiber.Router) *ioc.Container {
	applicationContainer := ioc.NewContainer(app, map[string]interface{}{
		"role": role.RoleModule(app),
	}, map[string]func() ioc.BaseObject{}, map[string]func() ioc.BaseRouter{}, map[string]func() ioc.BaseObject{})

	return applicationContainer
}
