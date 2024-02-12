package presenters

import (
	"taufikRahadi/fiber-boilerplate/application/role/ucase"
	"taufikRahadi/fiber-boilerplate/common/ioc"

	"github.com/gofiber/fiber/v2"
)

type RolePresenter struct {
	module *ioc.Container
}

func NewRolePresenter(module *ioc.Container) *RolePresenter {
	return &RolePresenter{
		module: module,
	}
}

func (r *RolePresenter) roleService() *ucase.RoleService {
	roleService, _ := r.module.CallProviders("RoleService", ucase.RoleService{})

	converted := roleService.(ucase.RoleService)
	return &converted
}

func (r *RolePresenter) GetUserList(c *fiber.Ctx) error {
	return c.JSON(r.roleService().ListRole())
}
