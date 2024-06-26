package bondApi

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"` // user-level status message
	//AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrRender(err error, code int, status string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: code,
		StatusText:     status,
		ErrorText:      err.Error(),
	}
}

const (
	ErrBasicAuth      = "failed basic auth"
	ErrNotFound       = "resource not found"
	ErrBadRequest     = "bad request"
	ErrNoAuth         = "no authorization"
	ErrInternalServer = "internal server error"
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
