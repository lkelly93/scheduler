// +build !longTests

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
	container "github.com/lkelly93/scheduler/internal/server_container"
)

func TestStartNewScheduler(t *testing.T) {
	schedulerName := "scheduler1"
	addr, err := container.StartNewScheduler(schedulerName)
	defer cleanupContainer(schedulerName, t)
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
