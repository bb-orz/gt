package libs

import (
	"fmt"
	"gt/utils"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

const SampleGitUrl = "https://github.com/bb-orz/goapp-sample.git"   // 最简模板
const AccountGitUrl = "https://github.com/bb-orz/goapp-account.git" // 通用账户模板

func GitCloneSample(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + SampleGitUrl + " " + appName + "`"
	err := ExecShellCommand(shellCmd)
	return err
}

func GitCloneAccount(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + AccountGitUrl + " " + appName + "`"
	err := ExecShellCommand(shellCmd)
	return err
}

const ShellToUse = "bash"

func ExecShellCommand(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}

func GetVer(v string) (uint64, uint64, error) {
	var err error
	var bigVer, subVer uint64
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	bigVer, err = strconv.ParseUint(v[2:first], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	if first == last {
		subVer, err = strconv.ParseUint(v[first+1:], 10, 64)
	} else {
		subVer, err = strconv.ParseUint(v[first+1:last], 10, 64)
	}

	return bigVer, subVer, err
}

func CheckGoMod() (string, bool) {
	var err error
	var version string
	var bigVer, subVer uint64

	version = runtime.Version()
	bigVer, subVer, err = GetVer(version)
	if err != nil {
		return version, false
	}
	if bigVer > 1 {
		// 未来的2.0版本大概率兼容
		return version, true
	}
	if subVer < 11 {
		// 子版本小于11不支持go mod
		return version, false
	} else if subVer < 13 && subVer >= 11 {
		// 子版本11、12需要开启环境变量GO111MODULE=on
		err = ExecShellCommand("export GO111MODULE=on")
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return version, true
		}
	}
	return version, true
}

// 替换项目中的主包名
func ReplaceMainPackageNAme(pwd, name string) error {
	utils.CommandLogger.Warning("Init", fmt.Sprintf("You create a new project,please replace main package name from 'goapp' to '%s'", name))
	// TODO 遍历项目每一个文件
	// TODO 读取每个文件前n行
	// TODO 替换成特定包名
	return nil
}
