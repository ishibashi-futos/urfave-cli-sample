package actions

import (
	"github.com/urfave/cli"

	"sample-cli/services"
)

func MainAction(context *cli.Context) error {
	var ss services.IService
	ss = services.NewMainService(context)
	ss.Run()
	return nil
}