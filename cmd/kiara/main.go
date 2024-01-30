package main

import (
	"fmt"
	"os"
	"github.com/Flavio-coutinho/container-kiara/pkg/container"
)

func main() {
	fmt.Println("Welcome to your container management App!")

	if len(os.Args) < 2 {
		fmt.Println("Usage: kiara <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "create":
			createContainer()
	case "start":
			startContainer()
	case "stop":
			stopContainer()
	case "delete":
			deleteContainer()
	default:
		fmt.Println("Invalid command. Usage: kiara <command>")
		os.Exit(1)
	}

	fmt.Println("Exiting the application")
}

func createContainer() {
	fmt.Println("Creating a new container")
	container.CreateContainer()
}

func startContainer() {
	fmt.Println("Starting a new container")
	container.StartingContainer()
}

func stopContainer() {
	fmt.Println("Stopping a container")
	container.StopContainer()
}

func deleteContainer() {
	fmt.Println("Deleting a container")
	container.DeleteContainer()
}