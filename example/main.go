package main

import (
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	"github.com/stevenxie/ticktick"
	ess "github.com/unixpickle/essentials"
)

func main() {
	c, err := ticktick.NewClient()
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
	if err := c.Login(user, string(pass)); err != nil {
		ess.Die("Failed to login:", err)
	}
	fmt.Print("Logged in successfully.\n\n")

	// Fetch tasks.
	fmt.Println("Fetching all remaining tasks...")
	tasks, err := c.ListTasks()
	if err != nil {
		ess.Die(err)
	}

	// Print tasks.
	fmt.Printf("Got %d tasks:\n", len(tasks))
	for _, task := range tasks {
		fmt.Printf("\t%s\n", task.Title)
	}
}
