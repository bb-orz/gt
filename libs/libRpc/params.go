package libRpc

type CmdParams struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	ProtoBufPath       string `json:"protubuf_path"`
	ClientOutputPath   string `json:"client_output_path"`
	ProtoGenOutputPath string `json:"proto_gen_output_path"`
	ServiceOutputPath  string `json:"service_output_path"`
	ServerOutputPath   string `json:"server_output_path"`
}
