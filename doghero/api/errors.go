package api

import "net/http"

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error       `json:"-"`
	HTTPStatusCode int         `json:"-"`
	StatusText     string      `json:"status"`
	AppCode        int64       `json:"code,omitempty"`
	Cause          string      `json:"error,omitempty"`
	Detail         interface{} `json:"detail"`
}

// NewHTTPError parse the given values to parse it to a commom error response
func NewHTTPError(content interface{}, status int, err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     http.StatusText(status),
		Cause:          err.Error(),
		Detail:         content,
	}
}
