// prove who you say who you are

package product

import (
	"fmt"
	"net/http"
)

func (product *ProductModel) Authentication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authentication in product module ")
}
