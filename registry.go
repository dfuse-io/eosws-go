package eosws

import "reflect"

var IncomingMessageMap = map[string]reflect.Type{}
var IncomingStructMap = map[reflect.Type]string{}
var OutgoingMessageMap = map[string]reflect.Type{}
var OutgoingStructMap = map[reflect.Type]string{}

func RegisterIncomingMessage(typeName string, obj interface{}) {
	refType := reflect.TypeOf(obj)
	IncomingMessageMap[typeName] = refType
	IncomingStructMap[refType] = typeName
}

func RegisterOutgoingMessage(typeName string, obj interface{}) {
	refType := reflect.TypeOf(obj)
	OutgoingMessageMap[typeName] = refType
	OutgoingStructMap[refType] = typeName
}

type OutgoingMessager interface {
	SetType(v string)
	SetReqID(v string)
}

// type IncomingMessager interface {
// 	GetType() string
// 	GetReqID() string
// 	GetCommon() CommonIn
// }
