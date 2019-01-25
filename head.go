package eosws

import "time"

func init()  {
	RegisterIncomingMessage("head_info", HeadInfo{})
	RegisterOutgoingMessage("get_head_info", GetHeadInfo{})
}

type HeadInfo struct {
	CommonIn
	Data struct{
		LastIrreversibleBlockNum	uint32		`json:"last_irreversible_block_num"`
		LastIrreversibleBlockId 	string		`json:"last_irreversible_block_id"`
		HeadBlockNum				uint32 		`json:"head_block_num"`
		HeadBlockID 				string 		`json:"head_block_id"`
		HeadBlockTime				time.Time	`json:"head_block_time"`
		HeadBlockProducer 			string 		`json:"head_block_producer"`
	} `json:"data"`
}

type GetHeadInfo struct {
	CommonOut
}