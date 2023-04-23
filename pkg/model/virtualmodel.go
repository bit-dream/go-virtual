package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

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
