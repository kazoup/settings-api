package handler

import (
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

//Restore struct
type Restore struct{}

//FromFile ...
func (p *Restore) FromFile(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Step1 ...
func (p *Restore) Step1(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Step2 ...
func (p *Restore) Step2(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Step3 ...
func (p *Restore) Step3(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Step4 ...
func (p *Restore) Step4(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}
