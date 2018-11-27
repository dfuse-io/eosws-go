package eosws

func init() {
	RegisterOutgoingMessage("unlisten", Unlisten{})
	RegisterIncomingMessage("unlistened", Unlistened{})
}

type Unlisten struct {
	CommonOut
	Data struct {
		ReqID string `json:"req_id"`
	} `json:"data"`
}

type Unlistened struct {
	CommonIn
	Data struct {
		Success bool `json:"success"`
	} `json:"data"`
}
