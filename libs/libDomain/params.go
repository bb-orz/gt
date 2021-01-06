package libDomain

/*命令参数*/
type CmdParams struct {
	Name        string `json:"name"`
	Driver      string `json:"driver"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	DbName      string `json:"database"`
	Table       string `json:"table"`
	User        string `json:"user"`
	Password    string `json:"password"`
	OutputPath  string `json:"output_path"`
	DOutputPath string `json:"dto_output_path"`
	Formatter   string `json:"formatter"`
}
