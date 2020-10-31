package user

import (
	"coffeeshop/plugins"
)

func CreateModelInst(plugin *plugins.Plugins) *UserModel {

	userModelInst := &UserModel{
		Routers: make([]map[string]interface{}, 0),
		plugin:  plugin,
	}

	routersConfig := []map[string]interface{}{
		{
			"Path":                  "/users",
			"Method":                "GET",
			"Handler":               userModelInst.Handler,
			"Authentication":        plugin.Authentication,
			"Authorization":         plugin.Authorization,
			"ParameterVerification": plugin.ParameterVerification,
		},
	}

	userModelInst.Routers = routersConfig

	return userModelInst
}
