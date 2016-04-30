package handler

import (
	"github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Ldap struct
type Ldap struct{}

//Create ...
func (p *Ldap) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Ldap) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Ldap) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Ldap) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
