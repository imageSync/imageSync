package src

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

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
}

//初始化配置文件
func initConfig() {
	viper.SetDefault("env", "default")
	viper.SetDefault("default.username", "admin")
	viper.SetDefault("default.password", "123456")
	viper.SetDefault("default.server_address", "registry.cn-shanghai.aliyuncs.com")
	viper.SetDefault("default.image_tag", "registry.cn-shanghai.aliyuncs.com/tay3223/images")

	//设定配置文件写入格式为json
	viper.SetConfigType("toml")

	//指定写入地址，且每一次都是覆盖式写入（因为用户每执行一次init子命令，此处就默认它已经做好了一切被覆盖的心理准备）
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	//往这个地方写入一个全局默认配置文件：~/.imageSync
	defaultConfigPath := homeDir + "/.imageSync"
	if err := viper.WriteConfigAs(defaultConfigPath); err != nil {
		return
	}
	fmt.Println("配置文件 ~/.imageSync 初始化成功...")
}
