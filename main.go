package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	controller "github.com/PrakPBP/Martini/controller"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/users", controller.GetAllUsers)
	m.Post("/users", binding.Bind(controller.User{}), controller.InsertUser)
	m.Put("/users/:id", binding.Bind(controller.User{}), controller.UpdateUser)
	m.Delete("/users/:id", controller.DeleteUser)

	m.RunOnAddr(":8080")
}
