package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"
)

// Panics if there is any error
func parse_response(resp *http.Response, err error) string {
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func get_entry(ctx *cli.Context) error {
	key := ctx.String("key")

	fmt.Printf("Get %s\n", key)

	url := fmt.Sprintf("http://localhost:3000/api/value/%s", key)

	resp, err := http.Get(url)

	fmt.Println("Data: ", parse_response(resp, err))

	return nil
}

func list_entries(ctx *cli.Context) error {
	url := "http://localhost:3000/api/values"

	resp, err := http.Get(url)

	fmt.Println("Data: ", parse_response(resp, err))

	return nil
}

func add_entry(ctx *cli.Context) error {
	key := ctx.String("key")
	value := ctx.String("value")

	data := url.Values{
		"key":   {key},
		"value": {value},
	}

	resp, err := http.PostForm("http://localhost:3000/api/value", data)

	parse_response(resp, err)
	return nil
}

func remove_entry(ctx *cli.Context) error {
	key := ctx.String("key")

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3000/api/value/%s", key), nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	parse_response(resp, err)

	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds an entry to the Key-Value server",
				Action:  add_entry,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true},
					&cli.StringFlag{Name: "value", Aliases: []string{"v"}, Required: true},
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rm"},
				Usage:   "Removes the specified entry from the Key-Value server, if it exists",
				Action:  remove_entry,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls"},
				Usage:   "Prints out all the existing Key-Value pairs",
				Action:  list_entries,
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Prints the given Key-Value pair, if it exists",
				Action:  get_entry,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
