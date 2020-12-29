package jsonhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", CreateUser)
	log.Println(http.ListenAndServe(":60001", nil))
}

// CreateUser handler
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	validationErr := validate(user)
	if validationErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    500,
			Message: validationErr.Error(),
		})
		return
	}

	user.ID = "1000000"
	json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: "OK",
		User:    &user,
	})
}
