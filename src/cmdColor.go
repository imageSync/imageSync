package src

import "runtime"

type Color struct {
	ReSet  string //重置颜色
	Red    string //红色
	Green  string //绿色
	Yellow string //黄色
	Blue   string //蓝色
	Purple string //紫色
	Cyan   string //青色
	Gray   string //灰色
	White  string //白色
}

func NewColor() Color {
	if runtime.GOOS == "windows" {
		return Color{
			ReSet:  "",
			Red:    "",
			Green:  "",
			Yellow: "",
			Blue:   "",
			Purple: "",
			Cyan:   "",
			Gray:   "",
			White:  "",
		}

	} else {
		return Color{
			ReSet:  "\033[0m",
			Red:    "\033[31m",
			Green:  "\033[32m",
			Yellow: "\033[33m",
			Blue:   "\033[34m",
			Purple: "\033[35m",
			Cyan:   "\033[36m",
			Gray:   "\033[37m",
			White:  "\033[97m",
		}
	}
}
