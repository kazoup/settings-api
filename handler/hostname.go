package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Hostname struct
type Hostname struct{}

//Create ...
func (p *Hostname) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Hostname) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Hostname) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Hostname) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
