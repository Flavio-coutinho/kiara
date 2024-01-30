package main

import (
	"fmt"
	"os"
	"github.com/Flavio-coutinho/kiara/pkg/container"
)

func main() {
	fmt.Println("Welcome to your container management App!")

	if len(os.Args) < 2 {
		fmt.Println("Usage: kiara <command> [containerID]")
		os.Exit(1)
	}

	containerManager := container.NewContainerManager()

	command := os.Args[1]
	switch command {
	case "create":
			createContainer(containerManager)
	case "start":
		if len(os.Args) < 3 {
			fmt.Println("Usage: kiara start [containerID]")
			os.Exit(1)
		}
			containerID := os.Args[2]
			startContainer(containerManager, containerID)
	case "stop":
			stopContainer(containerManager)
	case "delete":
			deleteContainer(containerManager)
	case "list":	
			listContainers(containerManager)
	default:
		fmt.Println("Invalid command. Usage: kiara <command>")
		os.Exit(1)
	}

	fmt.Println("Exiting the application")

}

func createContainer(containerManager *container.ContainerManager) {
	fmt.Println("Creating a new container")
	

	newContainer := containerManager.CreateContainer("example-container")

	if newContainer != nil {
		fmt.Printf("Created container %s (%s)\n", newContainer.Name, newContainer.ID)
	} else {
		fmt.Println("Failed to create container")
	}
}

func startContainer(containerManager *container.ContainerManager, containerID string) {
	fmt.Println("Starting a new container")

	container := containerManager.GetContainerByID(containerID)

	if container != nil {
		container.StartContainer()
		fmt.Println("Container started Successfully")
	} else {
		fmt.Println("Container not found")
	}
}

func stopContainer(containerManager *container.ContainerManager) {
	fmt.Println("Stopping a container")

	container := containerManager.GetContainerByID("id-do-container")
	container.StopContainer()
}

func deleteContainer(containerManager *container.ContainerManager) {
	fmt.Println("Deleting a container")
	
	container := containerManager.GetContainerByID("id-do-container")
	containerManager.DeleteContainer(container.ID)
}

func listContainers(containerManager *container.ContainerManager) {
	fmt.Println("Listing containers:")
	for _, c := range containerManager.GetContainers() {
		fmt.Printf("ID: %s\n", c.ID)
	}
}

