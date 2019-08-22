package main

import (
	"github.com/urfave/cli"
	"os"

	"sample-cli/actions"
)

func main() {
	app := cli.NewApp()

	app.Name = "sample-cli"
	app.Usage = "Self Service Portal Control CLI Tool."
	app.Version = "0.0.1"

	app.Action = actions.MainAction

	app.Run(os.Args)
}
