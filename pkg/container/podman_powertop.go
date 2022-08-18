package container

import "C"
import (
	"context"
	"fmt"
	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/containers/podman/v3/pkg/specgen"
	"log"
	"time"
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
func StartPowetopContainer() (entities.ContainerCreateResponse, error) {
	//creates a new OCI spec based on raw  image
	specGen := specgen.NewSpecGenerator(rawImage, true)

	//giving priledged permission to the contaner , important for powertop container
	specGen.Privileged = true

	specGen.Terminal = true

	//validating the spec
	err := specGen.Validate()
	if err != nil {
		log.Fatal("spec not valid %v", err)
	}

	//creating the  container
	r, err := containers.CreateWithSpec(*ctx, specGen, nil)

	if err != nil {
		log.Println(err)

	}

	// Container start
	fmt.Println("Starting PowerTop container.....")
	err = containers.Start(*ctx, r.ID, nil)
	if err != nil {
		log.Printf("Error in starting container %v", err)
		return r, err
	}
	return r, nil
}

func LoopContainer() error {
	r, err := StartPowetopContainer()
	if err != nil {
		return err
	}
	time.Sleep(3600 * time.Second)
	err = StopContainer(r)
	return err
}

func StopContainer(r entities.ContainerCreateResponse) error {
	err := containers.Stop(*ctx, r.ID, nil)
	return err
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
