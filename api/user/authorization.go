// prove you have been authorized to access ...

package user

import (
	"fmt"
	"net/http"
)

func (um *UserModel) Authorization(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authorization in user module ")
}
