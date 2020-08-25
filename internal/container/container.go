package container

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//Docker is a strucutre that represents a docker container
type Docker struct {
	Name string
}

//BuildImage rebuilds the docker image. This method takes a very long time to
//execute and should only be called at intial startup.
func BuildImage() {
	err := exec.Command("bash", "TarTheDocker.sh").Run()
	if err != nil {
		log.Fatal(err, " :unable to create tar of DockerFile")
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	buildCtx, err := os.Open("DockerTar.tar.gz")

	if err != nil {
		log.Fatal(err, ":unable to create tar")
	}

	suppressBuildOutput := true
	var tags []string = []string{"secure:latest"}

	buildOps := types.ImageBuildOptions{
		SuppressOutput: suppressBuildOutput,
		Dockerfile:     "Dockerfile",
		Tags:           tags,
		// BuildArgs:     {"-rm"},
	}

	buildResponse, err := cli.ImageBuild(
		context.Background(),
		buildCtx,
		buildOps)

	if err != nil {
		out, _ := exec.Command("ls").Output()
		log.Fatal(string(out) + err.Error())
		// log.Fatal(err, ": unable to build docker image")
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read the image build response")
	}
}
