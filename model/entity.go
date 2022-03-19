package model

import "github.com/google/uuid"

type JsonObject map[string]interface{}

type RequestMessageType int64

const (
	Create RequestMessageType = iota
)

type SocketRequest struct {
	Messages []SocketRequestMessage
}

type SocketRequestMessage struct {
	RequestType RequestMessageType
	Data        interface{}
}

type Entity struct {
	Id   uuid.UUID  `json:"id"`
	Tags []Tag      `json:"tags"`
	Data JsonObject `json:"data"`
}

type Tag struct {
	Id   uuid.UUID
	Data JsonObject
}