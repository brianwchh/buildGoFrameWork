// prove you have been authorized to access ...

package product

import (
	"fmt"
	"net/http"
)

func (product *ProductModel) Authorization(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authorization in product module ")
}
