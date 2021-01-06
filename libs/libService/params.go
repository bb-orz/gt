package libService

type CmdParams struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	IOutputPath string `json:"interface_output_path"`
	MOutputPath string `json:"implement_output_path"`
	DOutputPath string `json:"dto_output_path"`
}
