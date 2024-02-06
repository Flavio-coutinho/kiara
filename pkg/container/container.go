package container

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type ContainerManager struct {
	containers map[string]*Container
	mu         sync.Mutex
}

func NewContainerManager() *ContainerManager {
	return &ContainerManager{
		containers: make(map[string]*Container),
	}
}

type Container struct {
	ID    string
	Name  string
	State string
}

func (cm *ContainerManager) GetContainers() map[string]*Container {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	result := make(map[string]*Container, len(cm.containers))
	for key, value := range cm.containers {
			result[key] = value
	}

	fmt.Printf("GetContainers: %v\n", result)

	return result
}


func (cm *ContainerManager) CreateContainer(name string) (*Container, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	container := &Container{
			ID:    generateUniqueID(),
			Name:  name,
			State: "created",
	}

	if container == nil {
			return nil, fmt.Errorf("Failed to create container")
	}

	fmt.Printf("Created container %s (%s)\n", container.Name, container.ID)

	cm.containers[container.ID] = container
	fmt.Printf("Containers after creation: %v\n", cm.containers)

	return container, nil
}


func (c *Container) StartContainer() {
	fmt.Printf("Starting container %s (%s)...\n", c.Name, c.ID)

	if c.State == "running" {
		fmt.Println("Container is already running")
		return
	}

	fmt.Println("Initializing...")

	c.State = "running"
	fmt.Println("Container started successfully")
}

func (c *Container) StopContainer() {
	fmt.Printf("Stopping container %s (%s)...\n", c.Name, c.ID)
	c.State = "stopped"
}

func (cm *ContainerManager) DeleteContainer(containerID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	container, exists := cm.containers[containerID]
	if !exists {
		fmt.Printf("Container with ID %s not found\n", containerID)
		return
	}

	fmt.Printf("Deleting container %s (%s)...\n", container.Name, container.ID)
	delete(cm.containers, container.ID)
}

func (cm *ContainerManager) GetContainerByID(containerID string) (*Container, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	container, exists := cm.containers[containerID]
	fmt.Printf("GetContainerByID called for %s. Exists: %v\n", containerID, exists)

	if !exists {
			fmt.Printf("Containers in GetContainerByID: %v\n", cm.containers)
	}

	return container, exists
}


func generateUniqueID() string {
	return uuid.New().String()
}
