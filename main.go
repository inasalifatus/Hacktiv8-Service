package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var user = []User{
	{Name: "inas", Age: 25},
	{Name: "thalia", Age: 25},
}

var PORT = ":8091"

func main() {
	http.HandleFunc("/get-user", getusers)
	http.HandleFunc("/register", registeruser)

	fmt.Println("Application start on port", PORT)
	http.ListenAndServe(PORT, nil)

}

func registeruser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {

		// if want to input from FormValue

		// name := r.FormValue("Name")
		// age := r.FormValue("Age")
		// jsonAge, _ := strconv.Atoi(age)

		// newUser := User{
		// 	Name: name,
		// 	Age:  jsonAge,
		// }

		// user = append(user, newUser)
		// json.NewEncoder(w).Encode(newUser)

		//----------------------------------

		//if want to input value in JSON

		decoder := json.NewDecoder(r.Body)
		var newuserr User
		err := decoder.Decode(&newuserr)
		if err != nil {
			fmt.Println("error data")
		}

		user = append(user, newuserr)

		json.NewEncoder(w).Encode(newuserr)

		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
