package can_database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"go.einride.tech/can/pkg/dbc"
	"go.einride.tech/can/pkg/descriptor"
	"os"
	"strings"
)

type DbcData struct {
	Path     string
	Messages ecu_model.MessageMap
}

func MessageDefToMessage(def *dbc.MessageDef) descriptor.Message {
	signals := make([]*descriptor.Signal, 0)
	for _, signalDef := range def.Signals {
		signal := descriptor.Signal{
			Name:             string(signalDef.Name),
			Start:            uint8(signalDef.StartBit),
			Length:           uint8(signalDef.Size),
			IsBigEndian:      signalDef.IsBigEndian,
			IsSigned:         signalDef.IsSigned,
			IsMultiplexer:    signalDef.IsMultiplexerSwitch,
			IsMultiplexed:    signalDef.IsMultiplexed,
			MultiplexerValue: uint(signalDef.MultiplexerSwitch),
			Offset:           signalDef.Offset,
			Scale:            signalDef.Factor,
			Min:              signalDef.Minimum,
			Max:              signalDef.Maximum,
			Unit:             string(signalDef.Unit),
			Description:      string(""),
			// TODO: Implement the rest of the signal properties
		}
		signals = append(signals, &signal)
	}
	message := descriptor.Message{
		Name:       string(def.Name),
		ID:         uint32(def.MessageID),
		Signals:    signals,
		Length:     uint8(def.Size),
		IsExtended: dbc.MessageID(def.MessageID).IsExtended(),
	}
	return message
}

func UnmarshalMessagesFromDef(definitions []dbc.Def) []descriptor.Message {
	messages := make([]descriptor.Message, 0)

	for _, item := range definitions {
		if _, ok := item.(*dbc.MessageDef); ok {
			message := MessageDefToMessage(item.(*dbc.MessageDef))
			messages = append(messages, message)
		}
	}
	return messages

}

func MarshalMessagesToMap(messages []descriptor.Message) ecu_model.MessageMap {
	messageMap := make(ecu_model.MessageMap, 0)
	for _, message := range messages {
		messageMap[message.Name] = message
	}
	return messageMap
}

func LoadDbc(file string, dbcData *ecu_model.MessageMap) error {

	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error while reading the file: %d", err)
	}

	parser := dbc.NewParser(file, data)
	err = parser.Parse()
	if err != nil {
		return fmt.Errorf("error while parsing the database: %d", err)
	}

	messages := UnmarshalMessagesFromDef(parser.Defs())
	for _, message := range messages {
		_, ok := (*dbcData)[message.Name]
		if ok {
			return errors.New("duplicate message names detected, only unique message names across all dbc files permitted")
		}
		(*dbcData)[message.Name] = message
	}
	return nil
}

// LoadEcu Loads an individual .ecu file and modifies the existing model
func LoadEcu(file string) (*ecu_model.VirtualEcu, error) {
	if !strings.HasSuffix(file, ".ecu") {
		return nil, errors.New("file must have .ecu suffix")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading from file: %d", err)
	}

	ecu := ecu_model.VirtualEcu{}
	err = json.Unmarshal(data, &ecu)
	if err != nil {
		return nil, fmt.Errorf("error occurred while unmarshing file to ecu model: %d", err)
	}

	// Load dbc data
	dbcData := ecu_model.MessageMap{}
	for _, file := range ecu.Files {
		err := LoadDbc(file, &dbcData)
		if err != nil {
			return nil, fmt.Errorf("error parsing dbc file: %d", err)
		}
	}

	ecu.DbcData = &dbcData

	return &ecu, nil
}
