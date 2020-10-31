// prove who you say who you are

package plugins

import (
	"fmt"
	"net/http"
)

func (plugin *Plugins) Authentication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authentication from plugins' method ")
}
