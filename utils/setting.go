package utils

import (
	"io/ioutil"
	"os"
)

const SettingFile = ".setting.ini"

// 工具配置项操作记录器

func readSettingFile() ([]byte, error) {
	return ioutil.ReadFile(SettingFile)
}

// 初始化设置文件
func InitSetting(name, sample string) error {
	var err error
	var file *os.File
	// 检查设置文件是否存在
	file, err = os.Create(SettingFile)
	if err != nil {
		// 已存在直接返回
		if os.IsExist(err) {
			return nil
		}

		// 不够权限返回错误
		if os.IsPermission(err) {
			return err
		}
	}

	// 创建文件成功后，写入初始数据
	// 项目名称，
	_, err = file.WriteString("AppName:" + name)
	// 选择的模板名称
	_, err = file.WriteString("Sample:" + sample)

	return nil
}
