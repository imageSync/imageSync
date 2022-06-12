package src

import (
    "fmt"
    "log"
    "os/exec"
)

// MacDockerRunCheck 在Mac上检查docker服务是否运行
func MacDockerRunCheck() bool {
    /*
       被执行的shell命令为：
       launchctl list | grep com.docker.docker | awk '{print $1}'
    */
    cmd := exec.Command("bash", "-c", "launchctl list | grep com.docker.docker | awk '{print $1}'")
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }

    //fmt.Printf("打印Shell输出:\n%s\n", string(out))
    //fmt.Println(reflect.TypeOf(out))
    //fmt.Println(reflect.TypeOf(string(out)))

    if len(string(out)) == 0 {
        fmt.Println("检测到Docker daemon未启动...")
        return false
    } else {
        fmt.Println("Docker daemon运行中...")
        return true
    }
}
