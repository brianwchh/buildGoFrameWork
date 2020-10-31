package plugins

import (
	"net/http"
)

type Plugins struct {
	DB                       *DB
	Cache                    *Cache
	AuthenticationPtr        func(http.ResponseWriter, *http.Request)
	AuthorizationPtr         func(http.ResponseWriter, *http.Request)
	ParameterVerificationPtr func(http.ResponseWriter, *http.Request)
}

func CreatePlugins() *Plugins {
	// create an instance of Plugins, and initialize the function as empty function

	db := initDataBase()
	cache := initCache()

	return &Plugins{
		DB:                       db,
		Cache:                    cache,
		AuthenticationPtr:        func(w http.ResponseWriter, r *http.Request) {}, // empty func, incase crash...
		AuthorizationPtr:         func(w http.ResponseWriter, r *http.Request) {},
		ParameterVerificationPtr: func(w http.ResponseWriter, r *http.Request) {},
	}
}

func (plugin *Plugins) PluginDefers() {
	plugin.DB.db.Close()
	plugin.Cache.client.Close()
}
