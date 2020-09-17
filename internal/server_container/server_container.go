//Package container is a package that will handle the creation and running
//of the docker image that scheduler will run in
package server_container

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//StartNewScheduler starts a new scheduler with the given options.
//returns the IP address for the given scheduler.
func StartNewScheduler(schedulerName string) (string, error) {
	///Defaults
	dockerfile := "Dockerfile_standard"
	networkName := "scheduler-cluster"
	imageID := "lkelly93/scheduler_image:latest"

	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	err = createDefaultImageIfNeeded(
		cli,
		imageID,
		dockerfile)

	if err != nil {
		return "", err
	}

	err = createSchedulerClusterNetworkIfNeeded(cli, networkName)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{Image: imageID},
		&container.HostConfig{
			NetworkMode: container.NetworkMode(networkName),
			Privileged:  true,
		},
		nil,
		schedulerName,
	)

	if err != nil {
		return "", err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	//Get container IP
	info, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return "", err
	}

	IP := info.NetworkSettings.Networks[networkName].IPAddress
	works := make(chan bool)
	ctx, canelRoutine := context.WithCancel(context.Background())
	defer canelRoutine()

	go func(ctx context.Context) {
		requestBody, _ := json.Marshal(map[string]string{
			"Code": "print()",
		})
		select {
		case <-ctx.Done():
			return
		default:
			for {
				_, err := http.Post(
					fmt.Sprintf("http://%s:%d/execute/python", IP, 3000),
					"application/json",
					bytes.NewBuffer(requestBody),
				)
				if err == nil {
					works <- true
					return
				}
			}
		}
	}(ctx)

	timer := time.After(1 * time.Second)
	select {
	case <-works:
		return IP, nil
	case <-timer:
		return IP, &UnreachableContainerError{name: schedulerName}
	}
}

//createSchedulerClusterNetworkIfNeeded checks to see if the "Scheduler-cluser"
//network exists, if it does then it returns a NetworkMode that holds that
//network's info. If it doesn't then it creates it and returns the NetworkMode
//with the new networks info.
func createSchedulerClusterNetworkIfNeeded(cli *client.Client, networkName string) error {
	ctx := context.Background()
	networkMode := networkName

	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})

	for i := 0; i < len(networks); i++ {
		network := networks[i]
		if network.Name == networkMode {
			return nil
		}
	}

	_, err = cli.NetworkCreate(ctx, networkMode, types.NetworkCreate{})
	if err != nil {
		return err
	}

	return nil
}

func createDefaultImageIfNeeded(cli *client.Client, imageTag string, dockerfile string) error {
	//Check if Image exists and build it if not.
	imageExists, err := findImage(cli, imageTag)
	if err != nil {
		log.Fatal(err)
	}

	if !imageExists {
		pullDockerImageFromRepo(cli)
	}

	return nil
}

func pullDockerImageFromRepo(cli *client.Client) {
	ctx := context.Background()

	out, err := cli.ImagePull(
		ctx,
		"docker.io/lkelly93/scheduler_image:latest",
		types.ImagePullOptions{},
	)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(ioutil.Discard, out)
}

func findImage(cli *client.Client, imageTag string) (bool, error) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	for _, image := range images {
		if image.RepoTags[0] == imageTag {
			return true, nil
		}
	}

	return false, nil
}
