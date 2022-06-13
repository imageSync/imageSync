package src

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"strings"
)

// Pull 拉取docker镜像
func Pull(imageName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	fmt.Println("拉取镜像...")
	FormatOut(out)
}

// Push 推送docker镜像
func Push(imageName string) {
	readConfig()
	authConfig := types.AuthConfig{
		Username:      viper.GetString("username"),
		Password:      viper.GetString("password"),
		ServerAddress: viper.GetString("server_address"),
	}
	encodedJSON, _ := json.Marshal(authConfig)
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	imageNameTag := imageTagRename(imageName)
	if err := cli.ImageTag(context.Background(), imageName, imageNameTag); err != nil {
		return
	}
	out, err := cli.ImagePush(context.TODO(), imageNameTag, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		return
	}
	fmt.Println("\n\n\n\n\n\n推送镜像...")
	FormatOut(out)
	fmt.Printf("\n新的镜像地址为：" + imageNameTag + "\n\n")
}

func imageTagRename(imageName string) string {
	//切割字符串
	repoTagList := strings.Split(imageName, "/")
	//取切割后数组的最后一个
	repoTag := repoTagList[len(repoTagList)-1]
	//把:号替换为-号
	repoTag2 := strings.Replace(repoTag, ":", "-", -1)
	//把.号替换为-号
	//repoTag3 := strings.Replace(repoTag2, ".", "-", -1)
	//获取配置文件中的image_tag内容
	imageTag := viper.GetString("image_tag")
	//把:号替换为空
	imageTag2 := strings.Replace(imageTag, ":", "", -1)
	//拼接新的imageTag内容
	newImageTag := imageTag2 + ":" + repoTag2
	return newImageTag
}

func readConfig() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	defaultConfigPath := homeDir + "/.imageSync"
	viper.SetConfigFile(defaultConfigPath)
	//viper.AddConfigPath("$HOME/.imageSync")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败，请执行imagesync init命令，生成默认配置文件，并对内容进行修改...", err)
		return
	}
}
