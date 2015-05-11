package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandNew,
	commandImport,
	commandList,
	commandShow,
	commandDelete,
}

var commandNew = cli.Command{
	Name:  "new",
	Usage: "",
	Description: `
`,
	Action: doNew,
}

var commandImport = cli.Command{
	Name:  "import",
	Usage: "",
	Description: `
`,
	Action: doImport,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: `
`,
	Action: doList,
}

var commandShow = cli.Command{
	Name:  "show",
	Usage: "",
	Description: `
`,
	Action: doShow,
}

var commandDelete = cli.Command{
	Name:  "delete",
	Usage: "",
	Description: `
`,
	Action: doDelete,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doNew(c *cli.Context) {
}

func doImport(c *cli.Context) {
}

func doList(c *cli.Context) {
}

func doShow(c *cli.Context) {
}

func doDelete(c *cli.Context) {
}
