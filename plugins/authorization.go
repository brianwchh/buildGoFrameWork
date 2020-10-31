// prove you have been authorized to access ...

package plugins

import (
	"fmt"
	"net/http"
)

func (plugin *Plugins) Authorization(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authorization from plugins' method ")
}
