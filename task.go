package ticktick

import (
	"encoding/json"
	"fmt"

	ess "github.com/unixpickle/essentials"
)

// A Task is a thing that can be done.
type Task struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StartDate string `json:"startDate"`
	DueDate   string `json:"dueDate"`
	TimeZone  string `json:"timeZone"`
	IsAllDay  bool   `json:"isAllDay"`
	Priority  int8   `json:"priority"`
}

const (
	// listTasksURL is the URL used for batch listing remaining tasks.
	listTasksURL = baseURL + "/batch/check/0"
)

// ListTasks returns a lists all remaining (incomplete) TickTick tasks.
func (c *Client) ListTasks() (_ []Task, err error) {
	defer ess.AddCtxTo("ticktick", &err)

	res, err := c.HTTP.Get(listTasksURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // bad response
		return nil, errFromRes(res)
	}

	// Decode response body.
	var data struct {
		SyncTaskBean struct {
			Update []Task `json:"update"`
		} `json:"syncTaskBean"`
	}
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("couldn't decode response body: %v", err)
	}

	return data.SyncTaskBean.Update, nil
}
