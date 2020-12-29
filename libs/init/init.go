package init

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// 检查目录权限
func CheckDirMode() bool {
	// 获取当前目录
	dir, err := os.Getwd()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// 目录是否可读可写
	err = syscall.Access(dir, syscall.O_RDWR)
	return err == nil
}

const GitUrl = "https://github.com/bb-orz/goapp-sample.git"

func GitClone(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + GitUrl + " " + appName + "`"
	err := ExecShellCommand(shellCmd)
	return err
}

const ShellToUse = "bash"

func ExecShellCommand(command string) (error) {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}

func GetMinVer(v string) (uint64, error) {
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	if first == last {
		return strconv.ParseUint(v[first+1:], 10, 64)
	}
	return strconv.ParseUint(v[first+1:last], 10, 64)
}
