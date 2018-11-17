package ticktick

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		return fmt.Errorf("reading response body: %v", err)
	}
	defer res.Body.Close()

	return &APIError{
		StatusCode: res.StatusCode,
		Msg:        string(body),
	}
}
