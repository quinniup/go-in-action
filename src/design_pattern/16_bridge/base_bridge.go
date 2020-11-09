package _6_bridge

import "fmt"

// 桥接模式 分离抽象部分和实现部分，使得两部分可以独立拓展。
// 类似于策略模式，但区别在于策略模式封装一系列算法使得算法可以相互替换。

// 抽象单元
type AbstractMessage interface {
	SendMessage(text, target string)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{method: method}
}

func (c *CommonMessage) SendMessage(text, target string) {
	c.method.Send(text, target)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (c *UrgencyMessage) SendMessage(text, target string) {
	c.method.Send(text, target)
}

// 实现单元
type MessageImplementer interface {
	Send(text, target string)
}

type MessageSMS struct {
}

func InitSMS() *MessageSMS {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, target string) {
	fmt.Printf("send %s to %s via sms.\n", text, target)
}

type MessageEmail struct {
}

func InitEmail() *MessageEmail {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, target string) {
	fmt.Printf("send %s to %s via email.\n", text, target)
}
