package src

import (
	"context"
	"encoding/json"
	"fmt"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/wxnacy/wgo/arrays"
	"os"
	"strings"
)

// GetImageInfo 获取image的属性
func GetImageInfo() {
	imageName := "docker.io/alpine:3.12"
	repoTagList := strings.Split(imageName, "/")
	repoTag := repoTagList[len(repoTagList)-1]
	fmt.Println(repoTag)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, dockerTypes.ImageListOptions{})
	if err != nil {
		return
	}

	for _, image := range images {
		i := arrays.ContainsString(image.RepoTags, repoTag)
		if i >= 0 {
			fmt.Println(image.ID)
			fmt.Println(image.RepoDigests)
			fmt.Println(image.RepoTags)
			fmt.Println(image.VirtualSize)
			fmt.Println(image.Size)
			fmt.Println(image.SharedSize)
			fmt.Println(image.Created)
			fmt.Println(image.Containers)
			fmt.Println(image.Labels)
			fmt.Println(json.Marshal(image.ParentID))
		}
	}
}

// Exit 更优雅的退出进程
func Exit(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}
