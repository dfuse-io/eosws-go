package mdl

type TableOp struct {
	Operation   string `json:"op"`
	ActionIndex int    `json:"action_idx"`
	Payer       string `json:"payer"`
	Path        string `json:"path"`

	chunks []string
}
