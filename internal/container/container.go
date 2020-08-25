//Package container is a package that will handle the creation and running
//of the docker image that scheduler will run in
package container

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//Docker is a structure that represents a docker container
type Docker struct {
	Name string
}

//BuildImage rebuilds the docker image. This method takes a very long time to
//execute and should only be called at initial startup.
func BuildImage() {
	dockerfile := "Dockerfile"
	dockerTar := "DockerTar.tar.gz"

	err := zipDockerFile(dockerfile, dockerTar)
	if err != nil {
		log.Fatal(err, " :unable to create tar of DockerFile")
	}
	defer os.Remove(dockerTar)

	// err := exec.Command("bash", "TarTheDocker.sh").Run()
	// if err != nil {
	// 	log.Fatal(err, " :unable to create tar of DockerFile")
	// }

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	buildCtx, err := os.Open(dockerTar)
	if err != nil {
		log.Fatal(err, " :unable to open tar file")
	}

	suppressBuildOutput := true
	var tags []string = []string{"secure:latest"}
	var buildArgs = map[string]*string{
		"--rm": nil,
	}

	buildOps := types.ImageBuildOptions{
		SuppressOutput: suppressBuildOutput,
		Dockerfile:     dockerfile,
		Tags:           tags,
		BuildArgs:      buildArgs,
	}

	buildResponse, err := cli.ImageBuild(
		context.Background(),
		buildCtx,
		buildOps)

	if err != nil {
		log.Fatal(err, ": unable to build docker image")
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read the image build response")
	}
}
