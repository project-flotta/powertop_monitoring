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
	"github.com/project-flotta/powertop_container/pkg/stats"
	"log"
	"os"
	"time"
)

const (
	rawImage = "docker.io/sibseh/powertopcsv:t1"
	path     = "/var/tmp/powertop_report.csv"
	//MountPoint string
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
	ctx, err := bindings.NewConnection(
		context.Background(),
		"unix:/run/podman/podman.sock",
	)
	if err != nil {
		log.Printf(
			"cannot connect to podman :%v",
			err,
		)
	}
	return &ctx
}

func StartPowetopContainer() (entities.ContainerCreateResponse, error, string) {
	//creates a new OCI spec based on raw  image
	specGen := specgen.NewSpecGenerator(
		rawImage,
		false,
	)

	//giving priledged permission to the contaner , important for powertop container
	specGen.Privileged = true

	specGen.ContainerStorageConfig.Rootfs = "/var/tmp/"
	specGen.Terminal = true

	//validating the spec
	//err := specGen.Validate()
	//if err != nil {
	//	log.Printf(
	//		"spec not valid %v",
	//		err,
	//	)
	//}

	//creating the  container
	r, err := containers.CreateWithSpec(
		*ctx,
		specGen,
		nil,
	)
	MountPoint, err := containers.Mount(
		*ctx,
		r.ID,
		nil,
	)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	if err != nil {
		fmt.Println("error")
		log.Println(err)

	}

	// Container start
	fmt.Println("Starting PowerTop container.....")
	err = containers.Start(
		*ctx,
		r.ID,
		nil,
	)
	if err != nil {
		log.Printf(
			"Error in starting container %v",
			err,
		)
		return r, err, ""
	}
	return r, nil, MountPoint
}

func LoopContainer() (error, [][]string) {
	r, err, Mountpoint := StartPowetopContainer()
	data, err := stats.ReadCSV(Mountpoint + path)
	fmt.Println("done")
	if err != nil {
		return err, [][]string{}
	}
	//sleep for 1 hr
	time.Sleep(20 * time.Second)
	err = StopContainer(r)
	return err, data
}

func StopContainer(r entities.ContainerCreateResponse) error {
	err := containers.Stop(
		*ctx,
		r.ID,
		nil,
	)
	return err
}

func PullPowertopContainerImage() error {

	present, _ := images.Exists(
		*ctx,
		rawImage,
		nil,
	)
	if !present {
		return nil
	}
	fmt.Println("Pulling Powertop image...")
	_, err := images.Pull(
		*ctx,
		rawImage,
		nil,
	)
	if err != nil {
		log.Printf(
			"could not pull PowerTop Image %v",
			err,
		)
		return err
	}
	return nil
}
func ListImages() {
	imageSummary, err := images.List(
		*ctx,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var names []string
	for _, i := range imageSummary {
		names = append(
			names,
			i.RepoTags...,
		)
	}
	fmt.Println("Listing images...")
	fmt.Println(names)
}
