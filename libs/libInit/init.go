package libInit

import (
	"github.com/bb-orz/gt/utils"
	"runtime"
	"strconv"
	"strings"
)

const SampleGitUrl = "https://github.com/bb-orz/goapp-sample.git"   // 最简模板
const AccountGitUrl = "https://github.com/bb-orz/goapp-account.git" // 通用账户模板
const GrpcGitUrl = "https://github.com/bb-orz/goapp-grpc.git"       // grpc 应用
const MicroGitUrl = "https://github.com/bb-orz/goapp-micro.git"     // go-micro 应用

func GitCloneSample(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + SampleGitUrl + " " + appName + "`"
	err := utils.ExecShellCommand(shellCmd)
	return err
}

func GitCloneAccount(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + AccountGitUrl + " " + appName + "`"
	err := utils.ExecShellCommand(shellCmd)
	return err
}

func GitCloneGrpc(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + GrpcGitUrl + " " + appName + "`"
	err := utils.ExecShellCommand(shellCmd)
	return err
}

func GitCloneMicro(appName string) error {
	shellCmd := "`git clone -b master --single-branch --depth 1 " + MicroGitUrl + " " + appName + "`"
	err := utils.ExecShellCommand(shellCmd)
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
		err = utils.ExecShellCommand("export GO111MODULE=on")
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return version, true
		}
	}
	return version, true
}
