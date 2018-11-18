package ticktick

import (
	"encoding/json"

	ess "github.com/unixpickle/essentials"
)

const (
	// statusURL is the URL used for getting TickTick account info.
	statusURL = baseURL + "/user/status"
)

type accountInfo struct {
	InboxID string `json:"inboxId"`
}

// checkAccount polls the account status endpoint, updating the Client's
// internal account info.
func (c *Client) checkAccount() error {
	// Perform request.
	res, err := c.HTTP.Get(statusURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Decode response body.
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

	// Update Client.
	c.inboxID = info.InboxID
	return nil
}
