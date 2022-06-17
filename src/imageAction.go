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

var user User

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	defaultConfigPath := homeDir + "/.imageSync"
	viper.SetConfigFile(defaultConfigPath)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败，请执行imagesync init命令，生成默认配置文件，并对内容进行修改...", err)
		return
	}

	env := viper.GetString("env")
	user = NewUser(WithUsername(env+".username"), WithPassword(env+".password"), WithServerAddress(env+".server_address"), WithImageTag(env+".image_tag"))
}

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
	FormatOut(out, "pull")
}

// Push 推送docker镜像
func Push(imageName string) {
	//实例化一个Docker认证客户端
	authConfig := types.AuthConfig{
		Username:      user.Username,
		Password:      user.Password,
		ServerAddress: user.ServerAddress,
	}
	encodedJSON, _ := json.Marshal(authConfig)                //结构体转换为json
	authStr := base64.URLEncoding.EncodeToString(encodedJSON) //进行Base64编码

	//实例化一个docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//docker image tag
	imageNameTag, _ := imageTagRename(imageName)
	if err := cli.ImageTag(context.Background(), imageName, imageNameTag); err != nil {
		return
	}

	//push image
	out, err := cli.ImagePush(context.TODO(), imageNameTag, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		return
	}
	fmt.Println("\n\n\n\n\n\n推送镜像...")
	FormatOut(out, "push")

	//判断imageTag是否是以阿里云公有仓库地址为开头
	if strings.HasPrefix(imageNameTag, "registry.cn-shanghai.aliyuncs.com") {
		urlList := strings.Split(imageNameTag, ".cn-shanghai.aliyuncs.com")
		urlList[0] = urlList[0] + "-vpc.cn-shanghai.aliyuncs.com"
		urlList2 := strings.Join(urlList, "")
		fmt.Printf("\n新的镜像地址为（公网）：" + imageNameTag + "\n")
		fmt.Printf("新的镜像地址为（VPC）：" + urlList2 + "\n\n")
	} else {
		fmt.Printf("\n新的镜像地址为（公网）：" + imageNameTag + "\n\n")
	}

}

func imageTagRename(imageName string) (string, string) {
	//切割字符串
	repoTagList := strings.Split(imageName, "/")

	//取切割后数组的最后一个
	repoTag := repoTagList[len(repoTagList)-1]

	//把:号替换为-号
	repoTag2 := strings.Replace(repoTag, ":", "-", -1)
	//把.号替换为-号
	//repoTag3 := strings.Replace(repoTag2, ".", "-", -1)

	//获取配置文件中的image_tag内容
	imageTag := user.ImageTag

	//把:号替换为空
	imageTag2 := strings.Replace(imageTag, ":", "", -1)

	//拼接新的imageTag内容
	newImageTag := imageTag2 + ":" + repoTag2
	return newImageTag, repoTag2
}
