package handler

import (
	"github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Public struct
type Public struct{}

//Create ...
func (p *Public) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Public) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Public) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Public) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
