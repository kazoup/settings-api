package handler

import (
	"github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"net/http"
)

//User struct
type User struct{}

//Create ...
func (p *User) Create(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Update ...
func (p *User) Update(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}

//Read ...
func (p *User) Read(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement

	//MOCK START
	res.StatusCode = http.StatusOK

	/* Expects JSON obj:
	{
	  "url": string,
	  "id": int64,
	  "username": string,
	  "group_name": string,
	  "last_login": string,
	  "email": string,
	  "date_joined": string,
	  "created": string,
	  "updated": string,
	  "intercom": {
	    "created_at": int64,
	    "app_id": string,
	    "Appliance UUID": string,
	    "Build": string,
	    "email": string,
	    "user_hash": string,
	    "company": {
	      "plan": string,
	      "plan_started": int64,
	      "id": int64,
	      "name": string
	    }
	  }
	}
	*/
	ru := ResponseUser{}
	res.Body = ru.GetResponse()
	//MOCK END
	return nil
}

//Search ...
func (p *User) Search(ctx context.Context, req *api.Request,
	res *api.Response) error {
	//TODO Implement
	res.StatusCode = 200 //FIXME: use HTTP status code from the htp package
	res.Body = "TODO"
	return nil
}

//Delete ...
func (p *User) Delete(ctx context.Context, req *api.Request,
	res *api.Response) error {

	return nil
}
