// prove who you say who you are

package user

import (
	"fmt"
	"net/http"
)

func (um *UserModel) Authentication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authentication in user module ")
}
