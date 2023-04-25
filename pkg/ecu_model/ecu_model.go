package ecu_model

import (
	"go.einride.tech/can/pkg/descriptor"
)

type MessageMap map[string]descriptor.Message

// VirtualEcu main structure that defines a virtual ecu instance. Typically initialized from a .ecu file and loaded
type VirtualEcu struct {
	Name       string            `json:"name"`
	EcuChannel Channel           `json:"ecuChannel"`
	Files      []string          `json:"files"`
	Messages   *[]VirtualMessage `json:"messages,omitempty"`
	DbcData    *MessageMap       `json:"dbcData,omitempty"`
}

// Channel basic channel structure that defines how and at what rate the ecu will communicate at
type Channel struct {
	Interface      string          `json:"interface"`
	Channel        string          `json:"channel"`
	BaudRate       int             `json:"baudRate"`
	ChannelOptions *ChannelOptions `json:"channelOptions,omitempty"`
}

// ChannelOptions options for things such as packet spacing
type ChannelOptions struct {
	PacketSpacing *int `json:"packetSpacing,omitempty"` // time delta between each unique message transmission, in ms
}

// VirtualMessage defines what messages will be sent by the virtual ecu and how frequent
type VirtualMessage struct {
	Name                   string              `json:"name"`
	TransmissionOptions    TransmissionOptions `json:"transmissionOptions"`
	DefaultValueForSignals *int                `json:"defaultValueForSignals,omitempty"`
	Signals                *[]VirtualSignal    `json:"signals,omitempty"`
	MessageDefinition      *descriptor.Message `json:"messageDefinition,omitempty"`
}

// VirtualSignal signal that will be transmitted via a virtual message
type VirtualSignal struct {
	Name             string             `json:"name"`
	DefaultValue     int                `json:"defaultValue"`
	Value            *int               `json:"value,omitempty"`
	ValueQueue       *[]int             `json:"valueQueue,omitempty"`
	LoopQueue        *bool              `json:"loopQueue,omitempty"`
	ImportQueueFile  *string            `json:"importQueue,omitempty"`
	LastValueHold    *bool              `json:"lastValueHold,omitempty"`
	SignalDefinition *descriptor.Signal `json:"signalDefinition,omitempty"`
}

// SignalTrigger basic structure for a signal, as defined by a can id, that will trigger a message to be sent when
// received by the virtual ecu
type SignalTrigger struct {
	Id         uint32 `json:"id"`
	SignalName string `json:"signalName"`
}

// TransmissionOptions basic transmission options for a virtual ecu
type TransmissionOptions struct {
	TransmitRate           *int             `json:"transmitRate,omitempty"`
	TriggeredByMessageIds  *[]uint32        `json:"triggeredByMessageIds,omitempty"`
	TriggeredBySignalNames *[]SignalTrigger `json:"triggeredBySignalNames,omitempty"`
}

// ChannelMap maps the interface channel to a virtual message via dictionary
type ChannelMap map[string][]VirtualMessage
