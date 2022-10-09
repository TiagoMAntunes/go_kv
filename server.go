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
		return fiber.ErrInternalServerError
	}

	return ctx.Send(res)
}

func get_value(ctx *fiber.Ctx) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	key := ctx.Params("value")
	value, err := db.kv[key]
	if err != true {
		return fiber.ErrNotFound
	}

	return ctx.SendString(fmt.Sprintf("{ \"%s\": \"%s\" }", key, value))
}

type NewValue struct {
	Key   string `json:"key" xml:"key" form:"key"`
	Value string `json:"value" xml:"value" form:"value"`
}

func post_value(ctx *fiber.Ctx) error {
	v := new(NewValue)

	if err := ctx.BodyParser(v); err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.kv[v.Key] = v.Value

	res, err := json.Marshal(v)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Send(res)
}

func delete_value(ctx *fiber.Ctx) error {
	key := ctx.Params("value")

	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.kv, key)

	return nil
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

	app.Mount("/api/", api)

	db = SafeMap{kv: make(map[string]string)}

	app.Listen(":3000")
}
