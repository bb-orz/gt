package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
)

// 检查目录权限
func CheckDirMode() error {
	// 获取当前目录
	dir, err := os.Getwd()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	// 目录是否可读可写
	return syscall.Access(dir, syscall.O_RDWR)

}

const ShellToUse = "bash"

func ExecShellCommand(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}

// 获取路径最里目录
func GetLastPath(p string) string {
	base := path.Base(p)
	return base
}

// TODO 替换项目中的主包名
func ReplaceMainPackageNAme(pwd, name string) error {
	CommandLogger.Warning("Init", fmt.Sprintf("You create a new project,please replace main package name from 'goapp' to '%s'", name))
	// 遍历项目每一个文件
	// 读取每个文件前n行
	// 替换成特定包名
	return nil
}
