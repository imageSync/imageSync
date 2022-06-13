package src

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/wxnacy/wgo/arrays"
	"io"
	"os"
	"strings"
)

// FormatOut 格式化输出
func FormatOut(out io.Reader) {
	//方式一
	//io.Copy(os.Stdout, out)

	//方式二
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(out)
	//fmt.Println(buf)

	//方式三
	for {
		oneLine, err := bufio.NewReader(out).ReadBytes('\n')
		if err != nil {
			break
		}
		tmpMap := make(map[string]interface{})
		json.Unmarshal(oneLine, &tmpMap)
		fmt.Println(tmpMap)
	}
}

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

func Exit(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}
