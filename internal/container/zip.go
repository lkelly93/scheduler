package container

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func zipDockerFile(fileName string, outName string) error {

	//Open File to write tar to
	out, err := os.Create(outName)
	if err != nil {
		log.Fatal(err, " :unable to create tar file.")
	}
	defer out.Close()

	//Create tar writer
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	//Open DockerFile to compress
	dockerfile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer dockerfile.Close()

	info, err := dockerfile.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	// Use full path as name, if we don't do this the directory structure would
	// not be preserved
	header.Name = fileName

	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	//Write file to tar ball
	_, err = io.Copy(tw, dockerfile)
	if err != nil {
		return err
	}

	return nil
}
