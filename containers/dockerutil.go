package containers

/**
⊕ [go - Can I pin docker API version : client version 1.38 is too new. Maximum supported API version is 1.37 - Stack Overflow](https://stackoverflow.com/questions/51028784/can-i-pin-docker-api-version-client-version-1-38-is-too-new-maximum-supported)
⊕ [Examples using the Docker Engine SDKs and Docker API | Docker Documentation](https://docs.docker.com/develop/sdk/examples/#run-a-container)

*/

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

const (
	dockerClientVersion string = "1.37"
)

/*
PullImage ...
⊕ [Examples using the Docker Engine SDKs and Docker API | Docker Documentation](https://docs.docker.com/develop/sdk/examples/)

*/
func PullImage(imagePath string) {
	ctx := context.Background()
	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, imagePath, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)
}

// ListContainers list all containers
func ListContainers(displayImage bool) {
	ctx := context.Background()
	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		// fmt.Println(container.Image)
		if displayImage {
			fmt.Println(container.Image)
		} else {
			fmt.Println(container.Names)
		}
	}
}

// StopContainerByMatchImage stop specific container by match image name part
func StopContainerByMatchImage(imagePart string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {

		if strings.Contains(container.Image, imagePart) {

			fmt.Printf("\nWe found %s so lets stop it\n", imagePart)

			fmt.Print("Stopping container ", container.ID[:10], "... ")
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				panic(err)
			}
			fmt.Println("Success")
		}
	}
}

// RunContainer with imagePath -> alpine , args -> []string{"echo", "hello world"}
func RunContainer(imagePath string, args []string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
	if err != nil {
		panic(err)
	}

	fmt.Println(".. pull ", imagePath)
	reader, err := cli.ImagePull(ctx, imagePath, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imagePath,
		Cmd:   args,
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}

func ListImages() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
	if err != nil {
		panic(err)
	}
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
