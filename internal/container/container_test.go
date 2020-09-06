package container_test

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/lkelly93/scheduler/internal/container"
)

func TestBuildStandardImage(t *testing.T) {
	buildOptions := container.BuildImageOptions{
		Dockerfile: "Dockerfile_standard",
		Tags:       []string{"secure:latest"},
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

	foundSecure := false
	for _, image := range images {
		if image.RepoTags[0] == "secure:latest" {
			foundSecure = true
		}
	}
	if !foundSecure {
		t.Errorf("The image was not built")
	}
}

//TODO write tests for TarDockerFile
