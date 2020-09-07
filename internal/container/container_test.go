package container_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/lkelly93/scheduler/internal/container"
)

func TestBuildStandardImage(t *testing.T) {
	imageTag := "build_image_test_image:latest"
	buildOptions := container.BuildImageOptions{
		Dockerfile: "Dockerfile_standard",
		Tags:       []string{imageTag},
	}
	err := container.BuildImage(&buildOptions)
	if err != nil {
		t.Fatalf(err.Error())
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		t.Fatalf(err.Error())
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		t.Fatalf(err.Error())
	}

	foundImage := false
	for _, image := range images {
		if image.RepoTags[0] == imageTag {
			foundImage = true
		}
	}
	if !foundImage {
		t.Fatal("The image was not built")
	}

	//Cleanup test image.
	_, err = cli.ImageRemove(context.Background(), imageTag, types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})

	if err != nil {
		t.Errorf("Could not remove image, this may be because it was not created")
	}
}

func TestStartNewScheduler(t *testing.T) {
	opts := container.StartNewSchedulerOptions{
		ImageID:       "scheduler:latest",
		SchedulerName: "scheduler1",
	}

	addr, err := container.StartNewScheduler(&opts)
	// defer cleanupContainer(opts.SchedulerName, t)
	if err != nil {
		t.Error(err)
		t.Fatalf("Could not create container.")
	}

	//Check if server is up
	testCodeInput := "THIS IS A TEST"
	requestBody, err := json.Marshal(map[string]string{
		"Code": fmt.Sprintf("print(\"%s\")", testCodeInput),
	})

	if err != nil {
		t.Error(err)
		t.Fatal("Could not create request body")
	}

	response, err := http.Post(
		fmt.Sprintf("http://%s:%d/execute/python", addr, 3000),
		"application/json",
		bytes.NewBuffer(requestBody),
	)

	if err != nil {
		t.Error(err)
		t.Fatal("Could not connect to container")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		t.Fatal("Could not read in the response body")
	}

	bodyStr := string(body)
	testCodeOutput := fmt.Sprintf("{\"Stdout\":\"%s\\n\"}", testCodeInput)
	if bodyStr != testCodeOutput {
		t.Fatalf("Expected %s, but got %s", testCodeOutput, bodyStr)
	}
}

func cleanupContainer(SchedulerName string, t *testing.T) {
	t.Helper()
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		t.Error(err)
		t.Errorf("Could not create client for cleanup after test.")
	}

	err = cli.ContainerRemove(
		ctx,
		SchedulerName,
		types.ContainerRemoveOptions{Force: true})

	if err != nil {
		t.Error(err)
		t.Errorf("Could not remove %s", SchedulerName)
	}
}

//TODO write tests for TarDockerFile
