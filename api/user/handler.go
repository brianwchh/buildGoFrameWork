package user

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (um *UserModel) Handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type","application/json")
	// json.NewEncoder(w).Encode(books)

	// fmt.Println(r)

	db := um.plugin.DB

	type User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
		IsSuper  bool   `json:"isSuper"`
	}

	queries := `SELECT * 
		FROM User`

	insertRows, err := db.Query(queries)

	for insertRows.Next() {
		var user User

		elements := []interface{}{&user.ID, &user.Name, &user.Password, &user.Email, &user.IsSuper}
		// for each row scan the result into user object
		// err = insertRows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.IsSuper)
		err = insertRows.Scan(elements...)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("name : %s", user.Name)
	}

	insertRows.Close()

	// test cache
	cache := um.plugin.Cache
	err = cache.Set("sb", "10", 1*time.Minute)
	if err != nil {
		panic(err.Error())
	}

	val, errr := cache.Get("sb")
	if errr != nil {
		panic(errr.Error())
	}
	fmt.Printf("%s \n", val)

	fmt.Fprintln(w, "Hello user!")
}
