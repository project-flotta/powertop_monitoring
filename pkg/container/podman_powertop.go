package container

import (
	"context"
	"fmt"
	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
	"log"
	"os"
)

const (
	rawImage = "docker.io/sibseh/powertopimg"
)

var (
	ctx     *context.Context
	specGen specgen.SpecGenerator
)

func init() {
	ctx = StartingPodmanSocket()
}

//type PowerContainer struct{
//	image
//	state
//
//}

func StartingPodmanSocket() *context.Context {
	ctx, err := bindings.NewConnection(context.Background(), "unix:/run/podman/podman.sock")
	if err != nil {
		log.Printf("cannot connect to podman :%v", err)
	}
	return &ctx
}
func StartPowetopContainer() error {
	//pulling powerTop image from dockerhub
	err := PullPowertopContainerImage()
	if err != nil {
		return err
	}

	//creates a new OCI spec based on raw  image
	specGen := specgen.NewSpecGenerator(rawImage, true)

	//giving priledged permission to the contaner , important for powertop container
	specGen.Privileged = true

	specGen.Terminal = true

	//validating the spec
	err = specGen.Validate()
	if err != nil {
		log.Fatal("spec not valid %v", err)
	}

	//creating the  container
	r, err := containers.CreateWithSpec(*ctx, specGen, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container start
	fmt.Println("Starting PowerTop container.....")
	err = containers.Start(*ctx, r.ID, nil)
	if err != nil {
		log.Printf("Error in starting container %v", err)
		return err
	}
	return nil
}

func PullPowertopContainerImage() error {
	fmt.Println("Pulling Powertop image...")
	_, err := images.Pull(*ctx, rawImage, nil)
	if err != nil {
		log.Printf("could not pull PowerTop Image %v", err)
		return err
	}
	return nil
}
