package main

import (
	"context"
	"fmt"
)

func main() {
	app := InitializeApp()

	if err := app.Start(context.Background()); err != nil {
		fmt.Println("Error starting the application:", err)
	}
	go func() {
	}()
	<-app.Done()
}

// support migrations
// coding time
