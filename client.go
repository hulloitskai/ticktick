package ticktick

import (
	"net/http"
	"net/http/cookiejar"

	ess "github.com/unixpickle/essentials"
)

const (
	// baseURL is the root URL for the TickTick API.
	baseURL = "https://api.ticktick.com/api/v2"
)

// Client is capable of interacting with the TickTick API.
type Client struct {
	HTTP *http.Client   // HTTP client
	Jar  *cookiejar.Jar // cookie storage
}

// NewClient returns a new Client.
func NewClient() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, ess.AddCtx("ticktick: creating cookiejar", err)
	}

	client := &http.Client{Jar: jar}
	return &Client{HTTP: client, Jar: jar}, nil
}
