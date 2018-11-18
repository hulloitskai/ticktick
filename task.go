package ticktick

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	ess "github.com/unixpickle/essentials"
)

// A Task is a thing that can be done.
type Task struct {
	ID        string `json:"id,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	DueDate   string `json:"dueDate,omitempty"`
	TimeZone  string `json:"timeZone,omitempty"`
	IsAllDay  *bool  `json:"isAllDay,omitempty"`
	Priority  int8   `json:"priority"`
}

const (
	// listTasksURL is the URL used for batch listing remaining tasks.
	listTasksURL = baseURL + "/batch/check"

	// addTaskURL is the URL used for adding a new task.
	addTaskURL = baseURL + "/task"
)

// GetTasks returns a lists all remaining (incomplete) TickTick tasks.
func (c *Client) GetTasks() ([]*Task, error) {
	var (
		url      = fmt.Sprintf("%s/%d", listTasksURL, c.checkpoint)
		res, err = c.HTTP.Get(url)
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // bad response
		return nil, ess.AddCtx("ticktick", errFromRes(res))
	}

	// Decode response body.
	var (
		data struct {
			Checkpoint   uint64 `json:"checkPoint"`
			SyncTaskBean struct {
				Update []*Task `json:"update"`
			} `json:"syncTaskBean"`
		}
		dec = json.NewDecoder(res.Body)
	)
	if err = dec.Decode(&data); err != nil {
		return nil, ess.AddCtx("ticktick: decoding response body", err)
	}

	// Close response body.
	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	// Update internal checkpoint.
	c.checkpoint = data.Checkpoint
	c.updateCachedTasks(data.SyncTaskBean.Update)

	// Create array of tasks from task cache.
	tasks := make([]*Task, 0, len(c.tasks))
	for _, task := range c.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// AddTask adds a task to TickTick.
//
// Some important fields on 't' to fill out include t.Name and t.ProjectID.
func (c *Client) AddTask(t *Task) (updated *Task, err error) {
	// Encode task to JSON, and store in buf.
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(t); err != nil {
		return nil, ess.AddCtx("ticktick: encoding task", err)
	}

	// Create request.
	req, err := http.NewRequest("POST", addTaskURL, buf)
	if err != nil {
		return nil, ess.AddCtx("ticktick: creating request", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Perform request.
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // bad response
		return nil, ess.AddCtx("ticktick", errFromRes(res))
	}

	// Decode response body.
	updated = new(Task)
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(updated); err != nil {
		return nil, ess.AddCtx("ticktick: decoding response body", err)
	}

	// Close response body.
	if err = res.Body.Close(); err != nil {
		return nil, err
	}
	return updated, nil
}

/////////////////////////////
// EXPERIMENTAL / UNSTABLE
/////////////////////////////

// Required for batch adding tasks:
/*
	if c.inboxID == "" {
		if err := c.checkAccount(); err != nil {
			return err
		}
		if c.inboxID == "" { // checking the account revealed no inbox ID
			return errors.New("ticktick: could not determine inbox ID")
		}
	}
*/

// rateExp is the exponent that controls the rate of change between consecutive
// task IDs; each successive taskID is decreasing by a factor of approximately
// 10^(rateExp + 4).
const rateExp = 19

// nextTaskID generates a new task ID, given the previous task's ID.
//
// It creates an ID string that is less than the previous string, such that
// a task created with this ID will be on the top of the list.
func nextTaskID(prevID string) (string, error) {
	intID := new(big.Int)
	if _, err := fmt.Sscanf(prevID, "%x", intID); err != nil {
		return "", ess.AddCtx("scanning hex ID to big.Int", err)
	}

	// Maximum for random generation.
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(rateExp), nil)
	max := new(big.Int).Mul(exp, big.NewInt(9995))

	// Generate the random part of the ID
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", ess.AddCtx("generating random ID", err)
	}

	// Create ID diff.
	diff := new(big.Int).Mul(exp, big.NewInt(5))
	diff.Add(diff, random)

	// Create next ID.
	nextID := intID.Sub(intID, diff)
	return fmt.Sprintf("%x", nextID), nil
}
