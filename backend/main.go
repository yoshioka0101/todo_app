package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   	  int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	
	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})
	
	// Todoを作成するためのエンドポイント
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		// Bodyが空の時にJSONでエラーを出力する
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Satatus(201).JSON(todo)
	})

	// 更新をする
	app.Patch("api/todos/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Todosをスライスした内容から一致する値を取ってきている
		for i, todo := range todos {
			if fmt.Sptint(todo.ID) == id {
				todo[i].Comleted = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error":"Todo not found!"})
	})

}
