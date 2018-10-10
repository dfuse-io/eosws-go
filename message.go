package eosws

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CommonIn struct {
	Type  string `json:"type"`
	ReqID string `json:"req_id"`
}

type MsgIn struct {
	CommonIn
	Data json.RawMessage `json:"data"`
}

type CommonOut struct {
	Type         string `json:"type"`
	ReqID        string `json:"req_id,omitempty"`
	Fetch        bool   `json:"fetch,omitempty"`
	Listen       bool   `json:"listen,omitempty"`
	StartBlock   int64  `json:"start_block,omitempty"`
	WithProgress int64   `json:"with_progress,omitempty"`
}

func (c *CommonOut) SetType(v string)  { c.Type = v }
func (c *CommonOut) SetReqID(v string) { c.ReqID = v }

func setType(msg OutgoingMessager) {
	objType := reflect.TypeOf(msg).Elem()
	typeName := OutgoingStructMap[objType]
	if typeName == "" {
		panic(fmt.Sprintf("invalid or unregistered message type: %T", msg))
	}
	msg.SetType(typeName)
}

type MsgOut struct {
	CommonOut
	Data interface{}
}
