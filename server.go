package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type SafeMap struct {
	mu sync.RWMutex
	kv map[string]string
}

var db SafeMap

func get_all_values(ctx *fiber.Ctx) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	res, err := json.Marshal(db.kv)
	if err != nil {
		return err
	}

	return ctx.Send(res)
}

func get_value(ctx *fiber.Ctx) error {
	db.mu.RLock()
	db.mu.RUnlock()

	key := ctx.Params("value")
	value := db.kv[key]

	return ctx.SendString(fmt.Sprintf("{ \"%s\": \"%s\" }", key, value))
}

type NewValue struct {
	key   string `json:"key" xml:"key" form:"key"`
	value string `json:"value" xml:"value" form:"value"`
}

func post_value(ctx *fiber.Ctx) error {
	v := new(NewValue)

	if err := ctx.BodyParser(v); err != nil {
		return err
	}

	fmt.Printf("Data: %s %s\n", v.key, v.value)

	db.mu.Lock()
	db.mu.Unlock()

	db.kv[v.key] = v.value

	res, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return ctx.Send(res)
}

func delete_value(ctx *fiber.Ctx) error {
	return ctx.SendString("delete values!")
}

func update_value(ctx *fiber.Ctx) error {
	return ctx.SendString("update values!")
}

func main() {
	app := fiber.New()

	// Create routes
	api := fiber.New()

	// GET /api/values
	api.Get("/values", get_all_values)
	api.Get("/value/:value", get_value)
	api.Post("/value", post_value)
	api.Delete("/value/:value", delete_value)
	api.Put("/value/:value", update_value)

	app.Mount("/api/", api)

	db = SafeMap{kv: make(map[string]string)}

	app.Listen(":3000")
}
