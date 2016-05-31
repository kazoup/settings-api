package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//LDAP struct
type LDAP struct{}

//Create ...
func (p *LDAP) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *LDAP) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *LDAP) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *LDAP) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
