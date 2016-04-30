package handler

import (
	"github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Policy struct
type Policy struct{}

//Create ...
func (p *Policy) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *Policy) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *Policy) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Read ...
func (p *Policy) Search(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Policy) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
