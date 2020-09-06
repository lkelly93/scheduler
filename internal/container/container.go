//Package container is a package that will handle the creation and running
//of the docker image that scheduler will run in
package container

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//BuildImageOptions represents all the possible options for BuildImages
type BuildImageOptions struct {
	Dockerfile string
	Tags       []string
}

//BuildImage rebuilds the docker image. This method takes a very long time to
//execute and should only be called at startup.
func BuildImage(opts *BuildImageOptions) error {
	dockerfile := opts.Dockerfile
	tags := opts.Tags

	dockerTar := dockerfile + ".tar.gz"

	err := tarDockerFile(dockerfile, dockerTar)
	defer os.Remove(dockerTar)
	if err != nil {
		return err
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	buildCtx, err := os.Open(dockerTar)
	if err != nil {
		return err
	}

	suppressBuildOutput := true
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
		return err
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}
