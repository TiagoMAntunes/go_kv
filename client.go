package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func get_entry(ctx *cli.Context) error {
	key := ctx.String("key")

	fmt.Printf("Get %s\n", key)

	return nil
}

func list_entries(ctx *cli.Context) error {
	fmt.Println("List...")
	return nil
}

func add_entry(ctx *cli.Context) error {
	key := ctx.String("key")
	value := ctx.String("value")

	fmt.Printf("Add %s: %s\n", key, value)

	return nil
}

func remove_entry(ctx *cli.Context) error {
	key := ctx.String("key")
	value := ctx.String("value")

	fmt.Printf("Remove %s: %s\n", key, value)
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
					&cli.StringFlag{Name: "value", Aliases: []string{"v"}, Required: true},
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
