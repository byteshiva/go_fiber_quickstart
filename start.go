package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
)

type FruitInterface struct {
	Page   int
	Fruits []string
}

func main() {
	app := fiber.New()

	// Returns plain text.
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, Fiber!")
		// navigate to => http://localhost:3000/
	})

	app.Get("/", func(c *fiber.Ctx) {
		fmt.Println("1st route!")
		c.Append("Link", "http://google.com", "http://localhost")
		// => Link: http://localhost, http://google.com
		c.Append("Link", "Test")
		// => Link: http://localhost, http://google.com, Test
		c.Next()
	})

	// curl -X POST http://localhost:8080 -d user=john

	app.Post("/", func(c *fiber.Ctx) {
		// Get raw body from POST request:
		c.Body() // user=john
		fmt.Println(c.Body()) // user=john
	})

	app.Get("*", func(c *fiber.Ctx) {
		fmt.Println("2nd route!")
		c.Next(fmt.Errorf("Some error"))
	})

	app.Get("/", func(c *fiber.Ctx) {
		fmt.Println(c.Error()) // => "Some error"
		fmt.Println("3rd route!")
		c.Send("Hello, World!")
	})

	// Load static files like CSS, Images & JavaScript
	app.Static("/public", "./public")

	// Returns a local HTML file.
	app.Get("/hello", func(c *fiber.Ctx) {
		if err := c.SendFile("./templates/hello.html"); err != nil {
			c.Next(err)
		}
		// navigate to => http://localhost:3000/hello
	})

	// Use parameters
	app.Get("/parameter/:value", func(c *fiber.Ctx) {
		c.Send("Get request with value: " + c.Params("value"))
		// navigate to => http://localhost:3000/parameter/this_is_the_parameter
	})

	// GET /john
	app.Get("/:name/:age?", func(c *fiber.Ctx) {
		fmt.Printf("Name: %s, Age: %s", c.Params("name"), c.Params("age"))
		// => Name: john, Age:
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) {
		fmt.Printf("From: %s, To: %s", c.Params("from"), c.Params("to"))
		// => From: LAX, To: SFO
	})

	// Use wildcards to design your API.
	app.Get("/api/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*") + " -> do lookups with these values")
		// navigate to => http://localhost:3000/api/user/chris

		// return serialized JSON.
		if c.Params("*") == "fruits" {

			response := FruitInterface{
				Page:   1,
				Fruits: []string{"apple", "peach", "pear"},
			}

			responseJson, _ := json.Marshal(response)

			c.Send(responseJson)

			// navigate to => http://localhost:3000/api/fruits
		}
	})

	app.Listen(3000)
}
