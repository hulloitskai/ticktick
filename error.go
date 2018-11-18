package ticktick

import (
	"fmt"
	"io/ioutil"
	"net/http"

	ess "github.com/unixpickle/essentials"
)

// APIError is an error returned by a TickTick API call.
type APIError struct {
	StatusCode int    // status code of the response
	Msg        string // response body
}

func (err *APIError) Error() string {
	return fmt.Sprintf("%v (status code: %d)", err.Msg, err.StatusCode)
}

// errFromRes reads an error from an http.Response.
func errFromRes(res *http.Response) error {
	// Read response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ess.AddCtx("reading response body", err)
	}
	defer res.Body.Close()

	return &APIError{
		StatusCode: res.StatusCode,
		Msg:        string(body),
	}
}

// UnwrapAPIError unwraps an error into an APIError, if it was originally
// an APIError.
//
// If the original error is of some other form, UnwrapAPIError returns nil.
func UnwrapAPIError(err error) *APIError {
	for { // cyclical unwrapping of essentials.CtxErrors
		val, ok := err.(*ess.CtxError)
		if !ok {
			break
		}
		err = val.Original
	}

	// Return underlying APIError, or nil.
	switch val := err.(type) {
	case *APIError:
		return val
	default:
		return nil
	}
}
