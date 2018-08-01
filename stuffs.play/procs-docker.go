package main

import "github.com/samlet/gostuff/containers"

func main() {
	// containers.PullImage("docker.io/library/alpine")
	// containers.ListContainers(true)
	// containers.StopContainerByMatchImage("ubuntu")

	// alpine , args -> []string{"echo", "hello world"}
	// containers.RunContainer("alpine", []string{"echo", "hello world"})

	containers.ListImages()
}
