package src

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// NewCmdParams 命令行参数
func NewCmdParams() string {
	//定义子命令
	init := pflag.NewFlagSet("init", pflag.ExitOnError)

	//定义命令行参数
	image := pflag.StringP("image", "i", "", "海外的镜像地址（格式：docker.io/nginx:1.21.6）")
	env := pflag.StringP("env", "e", "", "选择环境，可以不填写（默认：default）")

	//如果系统参数少于2个，则退出程序
	if len(os.Args) < 2 {
		fmt.Println("缺少运行参数...")
		os.Exit(1)
	}

	//如果env参数长度为0，则强制env=default
	if len(*env) == 0 {
		viper.Set("env", "defautl")
	}

	//系统参数的第一位时main，判断第二位上的内容是什么，如果是init则执行对应的代码块，否则就去执行普通的命令行参数
	switch os.Args[1] {
	case "init":
		init.Parse(os.Args[2:])
		initConfig() //初始化配置文件
		return ""

	default:
		//接收命令行参数
		pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
		pflag.Usage = myUsage
		pflag.Parse()

		//如果没有输入命令行参数，则终止服务
		if len(*image) == 0 {
			fmt.Println("参数值不能为空... \033[1;31;8m（输入 --help 查看命令详细用法）\033[0m")
			os.Exit(1)
		}
		return *image
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

//pflg格式化输入
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

//命令行提示
func myUsage() {
	temp := `工具名称:  imageSync
工具版本:  0.0.2
工具描述:  加速拉取海外的docker镜像，并上传到自己的镜像仓库中。
详细文档:  https://github.com/tay3223/imageSync/blob/master/README.md


命令用法:  
1.运行 "imageSync init" 命令生成配置文件 ~/.imageSync
2.修改 "~/.imageSync" 配置文件中的内容
3.在终端中使用 imageSync 命令（例如：imageSync --help）


示例用法：
imageSync -e <env环境> -i <海外的镜像地址>
（或）
imageSync --env=<env环境> --image=<海外的镜像地址>


参数说明:
`
	fmt.Printf("\n------------------------\n")
	_, _ = fmt.Fprintf(os.Stderr, temp)
	pflag.PrintDefaults()
	fmt.Printf("------------------------\n\n")
}
