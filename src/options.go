package src

import (
    "fmt"
    "github.com/mitchellh/go-homedir"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
    "os"
    "strings"
)

// NewOptions 命令行参数
func NewOptions() string {
    //定义子命令
    init := pflag.NewFlagSet("init", pflag.ExitOnError)
    //initConfig := init.StringP("config", "c", "~/.imagesync", "配置文件路径")

    //定义命令行参数
    image := pflag.StringP("image", "i", "", "输入海外的镜像地址")
    //var image = pflag.StringP("image", "i", "docker.io/alpine:3.12", "镜像地址")

    //fmt.Println(*initConfig)
    //fmt.Println(*image)
    //fmt.Println(os.Args)

    //如果系统参数少于2个，则退出程序
    if len(os.Args) < 2 {
        fmt.Println("缺少运行参数...")
        os.Exit(1)
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

    uuu := `工具名称:  imageSync
工具版本:  0.0.2
工具描述:  可以借用docker运行时里配置好的代理，自动拉取海外的docker镜像，并上传到自己的镜像仓库中。

命令行用法:  imageSync -i <海外的镜像地址>

参数说明:
`

    fmt.Printf("\n------------------------\n")
    _, _ = fmt.Fprintf(os.Stderr, uuu)
    pflag.PrintDefaults()
    fmt.Printf("------------------------\n\n")

}

//初始化配置文件
func initConfig() {
    viper.SetDefault("username", "admin")
    viper.SetDefault("password", "123456")
    viper.SetDefault("server_address", "registry.cn-shanghai.aliyuncs.com")
    viper.SetDefault("image_tag", "registry.cn-shanghai.aliyuncs.com/tay3223/images")

    //设定配置文件写入格式为json
    viper.SetConfigType("json")

    //指定写入地址，且每一次都是覆盖式写入（因为用户每执行一次init子命令，此处就默认它已经做好了一切被覆盖的心理准备）
    homeDir, err := homedir.Dir()
    if err != nil {
        panic(err)
    }
    defaultConfigPath := homeDir + "/.imageSync"
    //err := viper.WriteConfigAs("~/.imageSync")

    if err := viper.WriteConfigAs(defaultConfigPath); err != nil {
        return
    }
    fmt.Println("配置文件初始化成功...")
}
