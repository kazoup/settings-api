package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//SmbUser struct
type SmbUser struct{}

//Create ...
func (p *SmbUser) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *SmbUser) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *SmbUser) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 501 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *SmbUser) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
