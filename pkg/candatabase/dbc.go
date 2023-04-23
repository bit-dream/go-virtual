package candatabase

import (
	"fmt"
	"go.einride.tech/can/pkg/dbc"
	"go.einride.tech/can/pkg/descriptor"
	"os"
)

type DbcData struct {
	Path     string
	Messages []descriptor.Message
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

func LoadDbc(file string) (DbcData, error) {
	dbcData := DbcData{Path: file}
	data, err := os.ReadFile(file)
	if err != nil {
		return dbcData, fmt.Errorf("error while reading the file: %d", err)
	}

	parser := dbc.NewParser(file, data)
	err = parser.Parse()
	if err != nil {
		return dbcData, fmt.Errorf("error while parsing the database: %d", err)
	}

	messages := UnmarshalMessagesFromDef(parser.Defs())
	dbcData.Messages = messages
	return dbcData, nil
}
