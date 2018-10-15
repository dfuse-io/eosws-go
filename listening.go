package eosws

func init() {
	RegisterIncomingMessage("listening", Listening{})
}

type Listening struct {
	CommonIn
}
