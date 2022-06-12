package main

import (
    "fmt"
    "imageSync/src"
    "os"
    "runtime"
)

func main() {

    //命令行参数（含子命令的实例化）
    imageName := src.NewOptions()
    if len(imageName) == 0 {
        os.Exit(33)
    }

    //判断当前系统类型
    switch runtime.GOOS {
    case "linux":
        fmt.Println("linux")

    case "windows":
        fmt.Println("windows")

    case "darwin":
        fmt.Println("当前操作系统为：MacOS...")

        //判断Docker daemon是否启动
        if src.MacDockerRunCheck() {
            src.Pull(imageName) //拉取镜像
            src.Push(imageName) //推送镜像
        } else {
            os.Exit(1)
        }

    default:
        fmt.Println("无法判断当前操作系统类型，程序停止运行....")
        os.Exit(1)
    }
}
