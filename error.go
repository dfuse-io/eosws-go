package eosws

func init() {
	RegisterIncomingMessage("error", Error{})
}

type Error struct {
	CommonIn

	Data struct {
		Code    string                 `json:"code"`
		Message string                 `json:"message"`
		Details map[string]interface{} `json:"details"`
	} `json:"data"`
}
