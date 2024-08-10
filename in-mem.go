package main

// app.Get("/api/todos", func(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(todos)
// })

// app.Post("/api/todos", func(c *fiber.Ctx) error {
// 	todo := &Todo{} //{id: 0, completed: false, body: """}

// 	if err := c.BodyParser(todo); err != nil {
// 		return err
// 	}

// 	if todo.Body == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required!"})
// 	}

// 	todo.ID = len(todos)

// 	todos = append(todos, *todo)

// 	return c.Status(201).JSON(todo)
// })

// // update a todo
// app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	for i, todo := range todos {
// 		if fmt.Sprint(todo.ID) == id {
// 			todos[i].Completed = true
// 			return c.Status(200).JSON(todos[i])
// 		}
// 	}

// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found!"})
// })

// // delete a todo
// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	for i, todo := range todos {
// 		if fmt.Sprint(todo.ID) == id {
// 			todos = append(todos[:i], todos[i+1:]...)
// 			return c.Status(200).JSON(fiber.Map{"success": "Todo successfully deleted!"})
// 		}
// 	}

// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found!"})
// })
