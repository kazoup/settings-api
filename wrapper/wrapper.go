package wrapper

import (
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"
)

// Validates user token
func ValidateUser(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// TODO: implement token validation.

		//START MOCK
		md, ok := metadata.FromContext(ctx)

		if !ok {
			return errors.InternalServerError("xx.yy.auth.wrapper", "Error retrieving request context")
		}

		if len(md["Authorization"]) > 0 {
			return fn(ctx, req, rsp)
		} else {
			return errors.Unauthorized("xx.yy.auth.wrapper", "Invalid token")
		}
		//END MOCK
	}
}
