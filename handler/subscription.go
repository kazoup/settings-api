package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Subscription struct
type Subscription struct{}

//Create ...
func (p *Subscription) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Subscription) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Subscription) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Subscription) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
