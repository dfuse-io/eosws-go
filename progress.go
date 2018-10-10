package eosws

func init() {
	RegisterIncomingMessage("progress", Progress{})
}

type Progress struct {
	CommonIn
	Data struct {
		BlockNum uint32 `json:"block_num"`
		BlockID  string `json:"block_id"`
	} `json:"data"`
}
