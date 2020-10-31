package user

import (
	"coffeeshop/plugins"
)

type UserModel struct {
	Routers []map[string]interface{}
	plugin  *plugins.Plugins
}

func (um *UserModel) GetRouters() []map[string]interface{} {
	return um.Routers
}
