package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Backup struct
type Backup struct{}

//Create ...
func (p *Backup) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Export ...
func (p *Backup) Export(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Search ...
func (p *Backup) Search(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Read ...
func (p *Backup) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *Backup) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
