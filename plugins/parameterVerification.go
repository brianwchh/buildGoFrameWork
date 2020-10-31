package plugins

import (
	"fmt"
	"net/http"
)

func (plugin *Plugins) ParameterVerification(w http.ResponseWriter, r *http.Request) {
	fmt.Println("parameter verification from plugins' method ")
}
