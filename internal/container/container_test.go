package container_test

import (
	"context"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/lkelly93/scheduler/internal/container"
)

func TestBuildImage(t *testing.T) {
	container.BuildImage()

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		log.Fatal(err, " :unable to get a list of all images in test method")
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
