package ticktick

import (
	"bytes"
	"fmt"
	"net/http"

	ess "github.com/unixpickle/essentials"
)

const (
	// loginURL is the URL used for authenticating with the TickTick platform.
	loginURL = baseURL + "/user/signon"
)

// Login authenticates with the TickTick API, providing access to other API
// methods.
func (c *Client) Login(user, pass string) error {
	// Write login JSON into buffer.
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{\"username\": \"%s\", \"password\": \"%s\"}", user, pass)

	// Create POST request.
	req, err := http.NewRequest("POST", loginURL, buf)
	if err != nil {
		return ess.AddCtx("ticktick: creating req", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Set query params.
	q := req.URL.Query()
	q.Add("wc", "true")
	q.Add("remember", "true")
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTP.Do(req)
	if err != nil {
		return ess.AddCtx("ticktick", err)
	}

	if res.StatusCode != 200 {
		return ess.AddCtx("ticktick", errFromRes(res))
	}
	return nil
}
