package wrapper

import (
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"
	"strings"
)

// ValidateUser validates user token
func ValidateUser(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// TODO: implement token validation.

		//START MOCK
		f := fn(ctx, req, rsp)

		// Avoid validation for Public.XXX
		// Would be nice to have wrapper by handler instead of by service
		if strings.Contains(req.Method(), "Public") {
			return f
		}

		md, ok := metadata.FromContext(ctx)

		if !ok {
			return errors.InternalServerError("xx.yy.auth.wrapper", "Error retrieving request context")
		}

		if len(md["Authorization"]) > 0 {
			return f
		} else {
			return errors.Unauthorized("xx.yy.auth.wrapper", "Invalid token")
		}
		//END MOCK
	}
}
