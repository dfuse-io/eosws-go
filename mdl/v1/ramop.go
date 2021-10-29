package mdl

type RAMOp struct {
	Operation   string `json:"op"`
	Family      string `json:"family"`
	Action      string `json:"action"`
	ActionIndex int    `json:"action_idx"`
	Payer       string `json:"payer"`
	Delta       int64  `json:"delta"`
	Usage       uint64 `json:"usage"` // new usage
}
