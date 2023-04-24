package model

import (
	"go.einride.tech/can/pkg/descriptor"
)

type VirtualModel struct {
	Messages []VirtualMessage `json:"messages"`
	Channels []Channel        `json:"channels"`
	Ecus     *[]Ecu           `json:"ecus,omitempty"`
}

type Ecu struct {
	Name       string           `json:"name"`
	EcuChannel Channel          `json:"ecuChannel"`
	Messages   []VirtualMessage `json:"messages"`
}

type Channel struct {
	Interface      string          `json:"interface"`
	ChannelName    string          `json:"channelName"`
	BaudRate       int             `json:"baudRate"`
	ChannelOptions *ChannelOptions `json:"channelOptions,omitempty"`
}

type VirtualMessage struct {
	Name                   string              `json:"name"`
	Dbc                    string              `json:"dbc"`
	Channel                string              `json:"channel"`
	TransmissionOptions    TransmissionOptions `json:"transmissionOptions"`
	DefaultValueForSignals *int                `json:"defaultValueForSignals,omitempty"`
	Signals                *[]VirtualSignal    `json:"signals,omitempty"`
	MessageDefinition      *descriptor.Message `json:"messageDefinition,omitempty"`
}

type VirtualSignal struct {
	Name                string               `json:"name"`
	DefaultValue        int                  `json:"defaultValue"`
	Value               *int                 `json:"value,omitempty"`
	ValueQueue          *[]int               `json:"valueQueue,omitempty"`
	LoopQueue           *bool                `json:"loopQueue,omitempty"`
	ImportQueueFile     *string              `json:"importQueue,omitempty"`
	LastValueHold       *bool                `json:"lastValueHold,omitempty"`
	TransmissionOptions *TransmissionOptions `json:"transmissionOptions,omitempty"`
	SignalDefinition    *descriptor.Signal   `json:"signalDefinition,omitempty"`
}

type TransmissionOptions struct {
	TransmitRate           *int      `json:"transmitRate,omitempty"`
	TriggeredByMessageIds  *[]int    `json:"triggeredByMessageIds,omitempty"`
	TriggeredBySignalNames *[]string `json:"triggeredBySignalNames,omitempty"`
}

type TriggerOptions struct {
	TransmitSignals *[]VirtualSignal `json:"transmitSignals,omitempty"`
}

type ChannelOptions struct {
	PacketSpacing *int `json:"packetSpacing,omitempty"`
}

type ChannelMap map[string][]VirtualMessage
