package ucase

import (
	commondto "taufikRahadi/fiber-boilerplate/common/dto"
	"taufikRahadi/fiber-boilerplate/common/ioc"
	"taufikRahadi/fiber-boilerplate/common/serializers"

	"gorm.io/gorm"
)

type RoleService struct {
	Module *ioc.Container
}

func NewService() ioc.BaseObject {
	return RoleService{
		Module: nil,
	}
}

func (roleService RoleService) Init(module *ioc.Container) ioc.BaseObject {
	roleService.Module = module

	return roleService
}

func (r *RoleService) ListRole() *commondto.BaseResponseSingle {
	roles := []Role{}
	r.Module.Deps["DB"].(*gorm.DB).Find(&roles)

	response := serializers.BuildResponseSingle("200", "Success", roles)

	return response
}
