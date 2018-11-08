package eosws

func init() {
	RegisterIncomingMessage("listening", Listening{})
}

type Listening struct {
	CommonOut
	NextBlock uint32 `json:"data"`
}
