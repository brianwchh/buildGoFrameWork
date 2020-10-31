package plugins

import (
	"fmt"
	"net/http"
)

/**

basically it is a url-handler-wrapper function.

func url-handler-wrapper ( f func(resp, req)) *handlerFunc {


	1. parsing parameter

	2. authentication / validation

	3. attach database / cache to each handler call

	4. router-handler

	5. post-handler preccssing if needed, before sending data back to client


}

**/

func (plugin *Plugins) HandlerWrapper(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/*************************** parameter verification ******************************/
		plugin.paramsVerificationWrapper(w, r)

		/*************************** authentication **********************************/
		plugin.authenWrapper(w, r)

		/*************************** authorization ***********************************/
		plugin.authorWrapper(w, r)

		next(w, r)

		fmt.Println("hello post-plugin")

		/**********************add post processing here *********************************/

	})

}

func (plugin *Plugins) paramsVerificationWrapper(w http.ResponseWriter, r *http.Request) {

	plugin.ParameterVerificationPtr(w, r)
}

func (plugin *Plugins) authorWrapper(w http.ResponseWriter, r *http.Request) {

	plugin.AuthorizationPtr(w, r)
}

func (plugin *Plugins) authenWrapper(w http.ResponseWriter, r *http.Request) {

	/***********************authentication per router *****************************************/
	plugin.AuthenticationPtr(w, r)
}
