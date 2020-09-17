// +build !longTests

package server_container_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	container "github.com/lkelly93/scheduler/internal/server_container"
)

var serverIP string

func TestMain(m *testing.M) {
	schedulerName := "scheduler1"
	IP, err := container.StartNewScheduler(schedulerName)
	serverIP = IP
	if err != nil {
		if _, ok := err.(*container.UnreachableContainerError); ok {
			cleanupContainer(schedulerName)
			log.Fatal(err)
		}
		log.Println("Could not create container.")
		log.Fatal(err)
	}
	code := m.Run()
	cleanupContainer(schedulerName)
	os.Exit(code)
}

func TestNewScheduler(t *testing.T) {
	testCodeInput := "THIS IS A TEST"
	requestBody, err := json.Marshal(map[string]string{
		"Code": fmt.Sprintf("print(\"%s\")", testCodeInput),
	})

	if err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("http://%s:%d/execute/python", serverIP, 3000)
	contentType := "application/json"
	body := bytes.NewBuffer(requestBody)

	response := post(url, contentType, body, t)
	actual := parseOutput(response.Body, t)
	response.Body.Close()

	expected := fmt.Sprintf("{\"Stdout\":\"%s\\n\"}", testCodeInput)
	if actual != expected {
		t.Fatalf("Expected %s, but got %s", expected, actual)
	}
}

func parseOutput(message io.ReadCloser, t *testing.T) string {
	t.Helper()
	output, err := ioutil.ReadAll(message)
	if err != nil {
		t.Fatal(err)
	}
	return string(output)

}

func post(url string, contentType string, body io.Reader, t *testing.T) *http.Response {
	t.Helper()
	response, err := http.Post(url, contentType, body)
	if err != nil {
		t.Fatal(err)
	}
	return response
}

func cleanupContainer(SchedulerName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Println(err)
		log.Fatal("Could not create client for cleanup after test.")
	}

	err = cli.ContainerRemove(
		ctx,
		SchedulerName,
		types.ContainerRemoveOptions{Force: true})

	if err != nil {
		log.Println(err)
		log.Printf("Could not remove %s", SchedulerName)
	}
}
