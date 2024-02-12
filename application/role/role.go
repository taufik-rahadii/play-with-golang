package role

import (
	"taufikRahadi/fiber-boilerplate/application/role/routers"
	"taufikRahadi/fiber-boilerplate/application/role/ucase"
	"taufikRahadi/fiber-boilerplate/common/ioc"
	"taufikRahadi/fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func RoleModule(app fiber.Router) *ioc.Container {
	c := ioc.NewContainer(
		app,
		map[string]interface{}{
			"DB": database.DB,
		},
		map[string]func() ioc.BaseObject{
			"RoleService": ucase.NewService,
		},
		map[string]func() ioc.BaseRouter{
			"RoleV1Route": routers.NewRoleV1Route,
		},
		make(map[string]func() ioc.BaseObject),
	)

	ucase.InitModel()

	return c
}
