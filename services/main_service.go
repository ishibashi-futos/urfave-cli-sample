package services

import (
	"github.com/urfave/cli"
	"fmt"
)

type MainService struct {
	context *cli.Context
}

func NewMainService(context *cli.Context) *MainService {
	service := &MainService{context}
	return service
}

func (service *MainService) Run() error {
	fmt.Println(service.context.Args().Get(0))
	return nil
}