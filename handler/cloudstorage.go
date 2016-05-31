package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//CloudStorage struct
type CloudStorage struct{}

//Create ...
func (p *CloudStorage) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *CloudStorage) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *CloudStorage) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *CloudStorage) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
