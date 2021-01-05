package libDomain

import (
	"fmt"
	"gt/utils"
	"io"
)

func CreateDomainFile(cmdParams *CmdParams) (domainFile, daoFile, testFile io.Writer, err error) {
	var (
		domainFileName, daoFileName, testFileName string
	)

	domainFileName = cmdParams.OutputPath + cmdParams.Name + "/" + cmdParams.Name + "_domain.go"
	daoFileName = cmdParams.OutputPath + cmdParams.Name + "/" + cmdParams.Name + "_dao.go"
	testFileName = cmdParams.OutputPath + cmdParams.Name + "/" + cmdParams.Name + "_test.go"

	if domainFile, err = utils.CreateFile(domainFileName); err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
		return
	} else {
		utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Create %s Domain File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), domainFileName))
	}
	if daoFile, err = utils.CreateFile(daoFileName); err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
		return
	} else {
		utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Create %s Domain Dao File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), daoFileName))
	}

	if testFile, err = utils.CreateFile(testFileName); err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
		return
	} else {
		utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Create %s Domain Test File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), testFileName))
	}

	return
}
