package container

import (
	"fmt"
	"sync"
	"github.com/google/uuid"
)

type ContainerManager struct {
	containers map[string]*Container
	mu   sync.Mutex 
}

func NewContainerManager() *ContainerManager {
	return &ContainerManager {
		containers: make(map[string]*Container),
	}
}

type Container struct {
	ID string 
	Name string
	State string
}


// func NewContainerManager() *ContainerManager {
// 	return &ContainerManager{
// 		containers: make([]*Container, 0),
// 	}
// }

func (cm *ContainerManager) GetContainers() map[string]*Container {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.containers
}

func (cm *ContainerManager) CreateContainer(name string) *Container {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	container := &Container{
		ID: generateUniqueID(),
		Name: name,
		State: "created",
	}
	cm.containers[container.ID] = container
	return container
}

func (c *Container) StartContainer() {
	fmt.Printf("Starting container %s (%s)...\n", c.Name, c.ID)
	c.State = "running"
}

func (c *Container) StopContainer() {
	fmt.Printf("Stopping container %s (%s)...\n", c.Name, c.ID)
	c.State = "stopped"
}

func(cm *ContainerManager) DeleteContainer(containerID string) {
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

func (cm *ContainerManager) GetContainerByID(containerID string) *Container {
	
	for _, c := range cm.containers {
		if c.ID == containerID {
			return c
		}
	}
	return nil
}

func generateUniqueID() string {
	return uuid.New().String()
}

