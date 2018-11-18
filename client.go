package ticktick

import (
	"net/http"
	"net/http/cookiejar"
)

const (
	// baseURL is the root URL for the TickTick API.
	baseURL = "https://api.ticktick.com/api/v2"
)

// Client is capable of interacting with the TickTick API.
type Client struct {
	HTTP *http.Client   // HTTP client
	Jar  *cookiejar.Jar // cookie storage

	tasks      map[string]*Task // cached tasks
	checkpoint uint64           // checkpoint ID for incremental updates

	inboxID string
}

// NewClient returns a new Client.
func NewClient() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	var (
		client = &http.Client{Jar: jar}
		tasks  = make(map[string]*Task)
	)

	return &Client{HTTP: client, Jar: jar, tasks: tasks}, nil
}

// updateCachedTasks updates c.tasks with 't', an array of new tasks.
func (c *Client) updateCachedTasks(t []*Task) {
	for _, task := range t {
		c.tasks[task.ID] = task
	}
}
