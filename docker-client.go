package main

import (
	"fmt"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

func main() {
	//defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	//cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Println(c.ID)
	}

	image_options := types.ImageListOptions{All: true}
	images, err := cli.ImageList(context.Background(), image_options)
	if err != nil {
		panic(err)
	}

	for _, i := range images {
		fmt.Println(i.ID)
		fmt.Println(i.RepoTags)
	}
}
