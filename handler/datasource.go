package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//DataSource struct
type DataSource struct{}

//Create ...
func (p *DataSource) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *DataSource) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *DataSource) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Search ...
func (p *DataSource) Search(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Discover ...
func (p *DataSource) Discover(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Delete ...
func (p *DataSource) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
