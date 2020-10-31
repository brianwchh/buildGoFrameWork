package main

import (
	"coffeeshop/api/product"
	"coffeeshop/api/user"
	"coffeeshop/plugins"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

/*
 ***********************************************************************************
	add modules here
	[
		"package name" : var Routers = []map[string]interface{}
	]
***************************************************************************************
*/

type ModuleInterface interface {
	GetRouters() []map[string]interface{}
	Handler(http.ResponseWriter, *http.Request)
}

func collectModules(plugin *plugins.Plugins) []map[string]ModuleInterface {

	modules := []map[string]ModuleInterface{
		{
			"ModelInstance": user.CreateModelInst(plugin),
		},
		{
			"ModelInstance": product.CreateModelInst(plugin),
		},
	}

	return modules
}

/*
*********************************************************************************************
*********************************************************************************************
 */

func MergeRouters(r *mux.Router, plugin *plugins.Plugins) {

	modules := collectModules(plugin)

	for j := 0; j < len(modules); j++ {
		modelInstance := modules[j]["ModelInstance"]

		routers := modelInstance.GetRouters()
		for i := 0; i < len(routers); i++ {
			path := routers[i]["Path"].(string)

			fmt.Println(" path : ", path)

			method := routers[i]["Method"].(string)
			handler := routers[i]["Handler"].(func(http.ResponseWriter, *http.Request))
			if routers[i]["Authentication"] != nil {
				plugin.AuthenticationPtr = routers[i]["Authentication"].(func(http.ResponseWriter, *http.Request))
			}
			if routers[i]["Authorization"] != nil {
				plugin.AuthorizationPtr = routers[i]["Authorization"].(func(http.ResponseWriter, *http.Request))
			}
			if routers[i]["ParameterVerification"] != nil {
				plugin.ParameterVerificationPtr = routers[i]["ParameterVerification"].(func(http.ResponseWriter, *http.Request))
			}

			r.HandleFunc(path, plugin.HandlerWrapper(handler)).Methods(method)

		}
	}
}

/*
*********************************************************************************************
			main
*********************************************************************************************
*/

func main() {

	/* create plugins instance  */
	plugin := plugins.CreatePlugins()
	defer plugin.PluginDefers()

	r := mux.NewRouter()

	MergeRouters(r, plugin)

	fmt.Println("server running ... ")

	log.Fatal(http.ListenAndServe(":8080", r))

}
