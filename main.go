package main

import (
	"fmt"
	"net/http"
)

const htmlLogin = `
<html>
<body>
	<form action="/login" method="POST">
		<label for="username">Username:</label><br>
		<input type="text" id="username" name="username"><br>
		<label for="password">Password:</label><br>
		<input type="password" id="password" name="password"><br><br>
		<input type="submit" value="Submit">
	</form> 
</body>
</html>
`

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get the username and password from the form
		username := r.FormValue("username")
		password := r.FormValue("password")

		checker := checkLogin(username, password)
		// Check the username and password against a database or other method
		if checker {
			// If the login is successful, redirect the user to the dashboard page
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			// If the login is unsuccessful, display an error message
			fmt.Fprint(w, "Incorrect username or password. Please try again.")
		}
	} else {
		// If the request is not a POST request, display the login form
		fmt.Fprint(w, htmlLogin)
	}
}

func checkLogin(username string, password string) (bool, error) {
	var correct bool
	correct = true
	return correct, nil

}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
