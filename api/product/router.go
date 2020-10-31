package product

import (
	"coffeeshop/plugins"
)

func CreateModelInst(plugin *plugins.Plugins) *ProductModel {

	productModelInst := &ProductModel{
		Routers: make([]map[string]interface{}, 0),
		plugin:  plugin,
	}

	routersConfig := []map[string]interface{}{
		{
			"Path":           "/products",
			"Method":         "GET",
			"Handler":        productModelInst.Handler,
			"Authentication": productModelInst.Authentication,
			"Authorization":  productModelInst.Authorization,
		},
	}

	productModelInst.Routers = routersConfig

	return productModelInst
}
