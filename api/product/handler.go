package product

import (
	"fmt"
	"net/http"
)

func (product *ProductModel) Handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type","application/json")
	// json.NewEncoder(w).Encode(books)

	// fmt.Println(r)

	fmt.Fprintln(w, "Hello products!")
}
