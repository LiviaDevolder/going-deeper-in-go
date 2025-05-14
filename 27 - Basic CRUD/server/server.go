package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser insert an user in db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Failed to read request body"))
		return
	}

	var user user

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, email) values (?, ?)")

	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}

	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	insertedId, err := insert.LastInsertId()

	if err != nil {
		w.Write([]byte("Error getting inserted Id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created succesfully! id: %d", insertedId)))
}

// GetUsers find all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	lines, err := db.Query("SELECT * FROM users")

	if err != nil {
		w.Write([]byte("Error getting users"))
		return
	}

	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error converting users to json"))
		return
	}
}

// GetUser find one user
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting param to int"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	line, err := db.Query("SELECT * FROM users WHERE id = ?", ID)

	if err != nil {
		w.Write([]byte("Error searching for user"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scaning user"))
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error converting user to JSON"))
		return
	}
}

// UpdateUser updates an user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting param to int"))
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")

	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Error updating user"))
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser remove user from db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting param to int"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Error deleting user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
