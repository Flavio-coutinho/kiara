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

	fmt.Println("Containers before creation", containerManager.GetContainers())

	newContainer, err := containerManager.CreateContainer("DefaultContainerName")
	if err != nil {
			fmt.Printf("Failed to create container: %v\n", err)
			return
	}

	fmt.Printf("Created container %s (%s)\n", newContainer.Name, newContainer.ID)
}


func startContainer(containerManager *container.ContainerManager, containerID string) {
	fmt.Println("Starting a new container")

	containers := containerManager.GetContainers()
	container, exists := containerManager.GetContainerByID(containerID)

	fmt.Println("Containers:", containers)

	if exists {
			container.StartContainer()
			fmt.Println("Container started Successfully")
	} else {
			fmt.Println("Container not found")
	}
}


func stopContainer(containerManager *container.ContainerManager) {
	fmt.Println("Stopping a container")

	container, exists := containerManager.GetContainerByID("id-do-container")
	if exists {
		container.StopContainer()
		fmt.Println("Container stopped successfully")
	} else {
		fmt.Println("Container not found")
	}
	
}

func deleteContainer(containerManager *container.ContainerManager) {
	fmt.Println("Deleting a container")
	
	container, exists := containerManager.GetContainerByID("id-do-container")
	if exists {
		containerManager.DeleteContainer(container.ID)
		fmt.Println("Container deleted successfully")
	} else {
		fmt.Println("Container not found")
	}
}

func listContainers(containerManager *container.ContainerManager) {
	fmt.Println("Listing containers:")
	containers := containerManager.GetContainers()

	if len(containers) == 0 {
		fmt.Printf("No containers found.")
		return
	}

	for _, c := range containers {
		fmt.Printf("ID: %s\n", c.ID)
	}
}

