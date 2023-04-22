package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type VirtualModel struct {
	Messages []VirtualMessage `json:"messages"`
	Channels []Channel        `json:"channels"`
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
	DefaultValueForSignals *int                `json:"defaultValueForSignals,omitempty"`
	Signals                *[]VirtualSignal    `json:"signals,omitempty"`
	TransmissionOptions    TransmissionOptions `json:"transmissionOptions"`
}

type VirtualSignal struct {
	Name                string               `json:"name"`
	DefaultValue        int                  `json:"defaultValue"`
	ValueQueue          *[]int               `json:"valueQueue,omitempty"`
	TransmissionOptions *TransmissionOptions `json:"transmissionOptions,omitempty"`
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

func LoadModel(file string) (*VirtualModel, error) {

	if !strings.HasSuffix(file, ".network") {
		return nil, errors.New("file must have .network suffix")
	}

	modelJson, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading from file: %d", err)
	}
	var vm VirtualModel
	err = json.Unmarshal(modelJson, &vm)
	if err != nil {
		return nil, fmt.Errorf("error occurred while unmarshing file to virtualization model: %d", err)
	}
	return &vm, nil
}

func GetMessages(vm VirtualModel) []VirtualMessage {
	return vm.Messages
}

func GetSignalsFromMessage(vm VirtualModel, messageName string) []VirtualSignal {
	messages := GetMessages(vm)
	for _, message := range messages {
		if message.Name == messageName {
			return *message.Signals
		}
	}
	return nil
}

type ChannelMap map[string][]VirtualMessage

func GetMessagesByChannels(vm VirtualModel) ChannelMap {
	chMap := make(map[string][]VirtualMessage)

	messages := GetMessages(vm)
	for _, message := range messages {
		if _, ok := chMap[message.Channel]; ok {
			// key exists in the map
			currentValue := chMap[message.Channel]
			newSlice := append(currentValue, message)
			chMap[message.Channel] = newSlice
		} else {
			messages := make([]VirtualMessage, 0)
			newSlice := append(messages, message)
			chMap[message.Channel] = newSlice
		}
	}

	return chMap
}
