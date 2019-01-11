package mdl

import (
	"time"
)

type Block struct {
	ID           string       `json:"id"`
	Irreversible bool         `json:"irreversible"`
	Header       *BlockHeader `json:"header"`
}

type BlockHeader struct {
	Timestamp        time.Time         `json:"timestamp"`
	Producer         string            `json:"producer"`
	Confirmed        uint16            `json:"confirmed"`
	Previous         string            `json:"previous"`
	TransactionMRoot string            `json:"transaction_mroot"`
	ActionMRoot      string            `json:"action_mroot"`
	ScheduleVersion  uint32            `json:"schedule_version"`
	NewProducers     *ProducerSchedule `json:"new_producers" eos:"optional"`
}

type ProducerSchedule struct {
	Version   uint32        `json:"version"`
	Producers []ProducerKey `json:"producers"`
}
type ProducerKey struct {
	AccountName     string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}
