package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Logs struct
type Logs struct{}

//Create ...
func (p *Logs) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Logs) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Logs) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Logs) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
