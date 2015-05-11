package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ekm"
	app.Version = Version
	app.Usage = ""
	app.Author = "Masataka Kuwabara"
	app.Email = "p.ck.t22@gmail.com"
	app.Commands = Commands

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
