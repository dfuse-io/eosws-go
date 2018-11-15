package eosws

func init() {
	RegisterIncomingMessage("listening", Listening{})
}

type Listening struct {
	CommonOut
	Data struct {
		NextBlock uint32 `json:"next_block"`
	} `json:"data"`
}
