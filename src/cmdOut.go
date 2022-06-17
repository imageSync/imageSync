package src

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// FormatOut 格式化输出
func FormatOut(out io.Reader, contextType string) {
	//方式一
	//io.Copy(os.Stdout, out)

	//方式二
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(out)
	//fmt.Println(buf)

	//方式三（逐行操作）
	color := NewColor() //实例化一个颜色渲染对象
	for {
		oneLine, err := bufio.NewReader(out).ReadBytes('\n')
		if err != nil {
			break
		}

		tmpMap := make(map[string]interface{})
		if err := json.Unmarshal(oneLine, &tmpMap); err != nil {
			return
		}

		//根据传入参数，来判断应该为打印加载什么颜色
		var outResult string

		if contextType == "pull" {
			outResult = strings.Replace(string(oneLine), "\\u003e", ">", -1)
			fmt.Println(color.Red + outResult + color.ReSet)

		} else if contextType == "push" {
			outResult = strings.Replace(string(oneLine), "\\u003e", ">", -1)
			fmt.Println(color.Blue + outResult + color.ReSet)

		} else {
			fmt.Println(tmpMap)
		}
	}
}
