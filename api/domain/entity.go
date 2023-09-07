package domain

import "time"

type message struct {
	Body MessageBody
	SendUser SendUser
	SendAt SendAt
}

type MessageBody struct {
	Value string
}

type SendUser struct {
	Value string
}

type SendAt struct {
	Value time.Time
}