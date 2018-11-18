package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/howeyc/gopass"
	"github.com/stevenxie/ticktick"
	ess "github.com/unixpickle/essentials"
)

func main() {
	client, err := ticktick.NewClient()
	if err != nil {
		ess.Die("Couldn't create client:", err)
	}

	// Get login details.
	var user string
	fmt.Print("Enter your TickTick username: ")
	fmt.Scanf("%s", &user)

	fmt.Print("Enter your TickTick password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		if err == gopass.ErrInterrupted { // occurs during CTRl+C
			os.Exit(0)
		}
		ess.Die("Failed to get password:", err)
	}

	// Login to the API.
	fmt.Printf("Logging in as '%s'...\n", user)
	if err := client.Login(user, string(pass)); err != nil {
		ess.Die("Failed to login:", err)
	}
	fmt.Print("Logged in successfully.\n\n")

	// Fetch tasks.
	fmt.Println("Fetching all remaining tasks...")
	tasks, err := client.GetTasks()
	if err != nil {
		ess.Die(err)
	}

	// Print tasks.
	fmt.Printf("Got %d tasks:\n", len(tasks))
	for _, task := range tasks {
		fmt.Printf("\t%s\n", task.Title)
	}

	// Add a new task.
	task := &ticktick.Task{
		Title:   "Study for linear algebra :(",
		Content: "😓",
	}

	task, err = client.AddTask(task)
	if err != nil {
		ess.Die(err)
	}

	fmt.Println("\nAdded a new task:")
	spew.Config.Indent = "\t"
	spew.Dump(task)
}
