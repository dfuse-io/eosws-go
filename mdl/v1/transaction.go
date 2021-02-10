package mdl

import (
	"encoding/json"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

// TransactionTrace maps to a `transaction_trace` in `chain/trace.hpp`
type TransactionTrace struct {
	ID              string                       `json:"id,omitempty"`
	BlockNum        uint32                       `json:"block_num"`
	BlockTime       eos.BlockTimestamp           `json:"block_time"`
	ProducerBlockID string                       `json:"producer_block_id"`
	Receipt         eos.TransactionReceiptHeader `json:"receipt"`
	Elapsed         int64                        `json:"elapsed"`
	NetUsage        uint64                       `json:"net_usage"`
	Scheduled       bool                         `json:"scheduled"`
	ActionTraces    []*ActionTrace               `json:"action_traces"`
	FailedDTrxTrace *TransactionTrace            `json:"failed_dtrx_trace"`
	Except          json.RawMessage              `json:"except"`
}

// ActionReceipt corresponds to an `action_receipt` from `chain/action_receipt.hpp`
type ActionReceipt struct {
	Receiver       string            `json:"receiver"`
	Digest         string            `json:"act_digest"`
	GlobalSequence eos.Uint64        `json:"global_sequence"`
	RecvSequence   eos.Uint64        `json:"recv_sequence"`
	AuthSequence   []json.RawMessage `json:"auth_sequence"`
	CodeSequence   eos.Uint64        `json:"code_sequence"`
	ABISequence    eos.Uint64        `json:"abi_sequence"`
}

// BaseActionTrace corresponds to a `base_action_trace` from `chain/trace.hpp`
type BaseActionTrace struct {
	Receipt          ActionReceipt      `json:"receipt"`
	Action           eos.Action         `json:"act"`
	ContextFree      bool               `json:"context_free"`
	Elapsed          int64              `json:"elapsed"`
	Console          string             `json:"console"`
	TransactionID    string             `json:"trx_id"`
	BlockNum         uint32             `json:"block_num"`
	BlockTime        eos.BlockTimestamp `json:"block_time"`
	ProducerBlockID  *string            `json:"producer_block_id,omitempty"`
	AccountRAMDeltas []*AccountRAMDelta `json:"account_ram_deltas"`
	Except           json.RawMessage    `json:"except"`
}

// AccountRAMDelta corresponds to an `account_delta` from `chain/trace.hpp`
type AccountRAMDelta struct {
	Account eos.AccountName `json:"account"`
	Delta   int64           `json:"delta"`
}

// ActionTrace corresponds to an `action_trace` from `chain/trace.hpp`
type ActionTrace struct {
	ActionOrdinal                          uint32 `json:"action_ordinal"`
	CreatorActionOrdinal                   uint32 `json:"creator_action_ordinal"`
	ClosestUnnotifiedAncestorActionOrdinal uint32 `json:"closest_unnotified_ancestor_action_ordinal"`
	BaseActionTrace
	Receiver     eos.AccountName `json:"receiver,omitempty"`
	InlineTraces []*ActionTrace  `json:"inline_traces"`
}

type PermissionLevel struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type TransactionLifecycle struct {
	TransactionStatus           string                 `json:"transaction_status"`
	ID                          string                 `json:"id"`
	Transaction                 *eos.SignedTransaction `json:"transaction"`
	ExecutionTrace              *TransactionTrace      `json:"execution_trace"`
	ExecutionBlockHeader        *eos.BlockHeader       `json:"execution_block_header"`
	DTrxOps                     []*DTrxOp              `json:"dtrxops"`
	CreationTree                CreationFlatTree       `json:"creation_tree"`
	DBOps                       []*DBOp                `json:"dbops"`
	RAMOps                      []*RAMOp               `json:"ramops"`
	TableOps                    []*TableOp             `json:"tableops"`
	PubKeys                     []*ecc.PublicKey       `json:"pub_keys"`
	CreatedBy                   *ExtDTrxOp             `json:"created_by"`
	CanceledBy                  *ExtDTrxOp             `json:"canceled_by"`
	ExecutionIrreversible       bool                   `json:"execution_irreversible"`
	DTrxCreationIrreversible    bool                   `json:"dtrx_creation_irreversible"`
	DTrxCancelationIrreversible bool                   `json:"dtrx_cancelation_irreversible"`
}

type ActionRef struct {
	BlockID     string `json:"block_id"`
	TrxIndex    int    `json:"trx_index"`
	TrxID       string `json:"trx_id"`
	ActionIndex int    `json:"action_index"`
}

type CreationFlatTree = []CreationFlatNode
type CreationFlatNode = [3]int
