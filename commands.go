package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
)

var regionFlag = cli.StringFlag{
	Name:  "region",
	Value: "",
}

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
	Flags:  []cli.Flag{regionFlag},
	Action: setAWSConfig(doNew),
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func setAWSConfig(f func(*cli.Context, *aws.Config)) func(c *cli.Context) {
	return func(c *cli.Context) {
		region := c.String("region")
		cfg := &aws.Config{
			Region: region,
		}
		f(c, cfg)
	}
}

func doNew(ctx *cli.Context, cfg *aws.Config) {
	name := ctx.Args()[0]
	e := ec2.New(cfg)
	keyOut, err := e.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: toPtr(name),
	})
	assert(err)
	fmt.Print(*keyOut.KeyMaterial)
}

func doImport(c *cli.Context) {
}

func doList(c *cli.Context) {
}

func doShow(c *cli.Context) {
}

func doDelete(c *cli.Context) {
}

func toPtr(s string) *string {
	return &s
}
