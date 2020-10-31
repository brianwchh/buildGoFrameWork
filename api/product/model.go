package product

import (
	"coffeeshop/plugins"
)

type ProductModel struct {
	Routers []map[string]interface{}
	plugin  *plugins.Plugins
}

func (product *ProductModel) GetRouters() []map[string]interface{} {
	return product.Routers
}
