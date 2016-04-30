package main

import (
	"log"
	"time"

	"github.com/kazoup/settings-api/handler"
	"github.com/micro/go-micro"
)

func main() {
	//create service
	service := micro.NewService(
		//TODO change namespace to com.kazoup.api
		micro.Name("go.micro.api.v2.settings"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)
	//Init service
	service.Init()
	//Register public settings handler
	//http://localhost:8080/v2/settings/public/{create,update,read,delete}
	service.Server().Handle(
		service.Server().NewHandler(new(handler.Public)),
	)
	//Register hostname settings handler
	//http://localhost:8080/v2/settings/hostname/{create,update,read,delete}
	service.Server().Handle(
		service.Server().NewHandler(new(handler.Hostname)),
	)
	//Register network credentials settings handler
	//http://localhost:8080/v2/settings/smbuser/{create,update,read,delete}
	service.Server().Handle(
		service.Server().NewHandler(new(handler.SmbUser)),
	)
	//Register network credentials settings handler
	//http://localhost:8080/v2/settings/cloudstorage/{create,update,read,delete}
	service.Server().Handle(
		service.Server().NewHandler(new(handler.CloudStorage)),
	)
	//Register network credentials settings handler
	//http://localhost:8080/v2/settings/ldap/{create,update,read,delete}
	service.Server().Handle(
		service.Server().NewHandler(new(handler.Ldap)),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
