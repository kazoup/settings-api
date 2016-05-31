package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//SMTP struct
type SMTP struct{}

//Create ...
func (p *SMTP) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *SMTP) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *SMTP) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *SMTP) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
