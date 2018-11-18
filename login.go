package ticktick

import (
	"bytes"
	"encoding/json"
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

	// Create request.
	req, err := http.NewRequest("POST", loginURL, buf)
	if err != nil {
		return ess.AddCtx("ticktick: creating request", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Set query params.
	q := req.URL.Query()
	q.Add("wc", "true")
	q.Add("remember", "true")
	req.URL.RawQuery = q.Encode()

	// Perform request.
	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 { // bad response
		return ess.AddCtx("ticktick", errFromRes(res))
	}

	// Extract inbox ID from response.
	var (
		info accountInfo
		dec  = json.NewDecoder(res.Body)
	)
	if err = dec.Decode(&info); err != nil {
		return ess.AddCtx("ticktick: decoding response body", err)
	}

	// Close response body.
	if err = res.Body.Close(); err != nil {
		return err
	}

	// Update client.
	c.inboxID = info.InboxID
	return nil
}
