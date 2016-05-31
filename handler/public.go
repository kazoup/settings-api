package handler

import (
	"net/http"

	api "github.com/micro/micro/api/proto"
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

/*
Read ...

Expects JSON obj:
{
  "APPLIANCE_IS_REGISTERED": bool,
  "APPLIANCE_IS_CONFIGURED": bool,
  "APPLIANCE_IS_DEMO": bool,
  "GIT_COMMIT_STRING": string,
  "SMB_USER_EXISTS": bool
}
*/
func (p *Public) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement

	//START MOCK
	res.StatusCode = http.StatusOK
	rp := &ResponsePublic{}
	res.Body = rp.GetResponse()
	//END MOCK

	return nil
}

//Delete ...
func (p *Public) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
