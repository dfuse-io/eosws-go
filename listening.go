package eosws

func init() {
	RegisterIncomingMessage("listening", Progress{})
}

type Listening struct {
	CommonOut
	NextBlock uint32 `json:"data"`
}
