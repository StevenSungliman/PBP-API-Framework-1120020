package controller

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
)

func GetAllUsers(r render.Render) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")

	if err != nil {
		r.JSON(200, map[string]interface{}{
			"status":  400,
			"message": "Get all users failed",
		})
	} else {
		var users []User
		for rows.Next() {
			var user User
			rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Type)
			users = append(users, user)
		}
		r.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Success",
			"data":    users,
		})
	}
}

func InsertUser(r render.Render, params martini.Params, user User) {
	db := connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO user(name, email, password, type) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Type)

	if err != nil {
		r.JSON(200, map[string]interface{}{
			"status":  400,
			"message": "Insert user failed",
		})
	} else {
		r.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Success",
		})
	}
}

func UpdateUser(r render.Render, params martini.Params, user User) {
	db := connect()
	defer db.Close()

	rows, errQ := db.Query("SELECT * FROM user WHERE id=?", params["id"])

	if errQ != nil {
		r.JSON(200, map[string]interface{}{
			"status":  400,
			"message": "Get user by id failed",
		})
	} else {
		var userDefault User

		for rows.Next() {
			rows.Scan(&userDefault.Id, &userDefault.Name, &userDefault.Email, &userDefault.Password, &userDefault.Type)
		}

		if user.Name == "" {
			user.Name = userDefault.Name
		}
		if user.Email == "" {
			user.Email = userDefault.Email
		}
		if user.Password == "" {
			user.Password = userDefault.Password
		}
		if user.Type == 0 {
			user.Type = userDefault.Type
		}
	}

	_, err := db.Exec("UPDATE user SET name=?, email=?, password=?, type=? WHERE id=?", user.Name, user.Email, user.Password, user.Type, params["id"])

	if err != nil {
		r.JSON(200, map[string]interface{}{
			"status":  400,
			"message": "Update user failed",
		})
	} else {
		r.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Success",
		})
	}
}

func DeleteUser(r render.Render, params martini.Params) {
	db := connect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM user WHERE id=?", params["id"])

	if err != nil {
		r.JSON(200, map[string]interface{}{
			"status":  400,
			"message": "Delete user failed",
		})
	} else {
		r.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Success",
		})
	}
}
