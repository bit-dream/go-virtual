package candatabase_test

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/can_database"
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"go.einride.tech/can/pkg/descriptor"
	"testing"
)

func TestLoadDbcFile(t *testing.T) {
	dbcData := ecu_model.MessageMap{}
	err := loaders.LoadDbc("tesla_can.dbc", &dbcData)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadingOfMessages(t *testing.T) {
	data := ecu_model.MessageMap{}
	err := loaders.LoadDbc("tesla_can.dbc", &data)
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("no messages collected during parsing of dbc file")
	}
	if len(data) != 42 {
		t.Error("expected number of parsed messages from dbc file is incorrect, should have gotten 42")
	}
}

func TestCreatePayload(t *testing.T) {
	signal1 := ecu_model.VirtualSignal{
		Name:         "Test1",
		DefaultValue: 10,
		SignalDefinition: &descriptor.Signal{
			Name:        "Test1",
			Start:       0,
			Length:      8,
			IsBigEndian: false,
			IsSigned:    false,
		},
	}
	signal2 := ecu_model.VirtualSignal{
		Name:         "Test2",
		DefaultValue: 200,
		SignalDefinition: &descriptor.Signal{
			Name:        "Test1",
			Start:       8,
			Length:      8,
			IsBigEndian: false,
			IsSigned:    false,
		},
	}

	arr := make([]ecu_model.VirtualSignal, 0)
	arr = append(arr, signal1, signal2)
	payload := can_database.GeneratePayloadFromSignals(arr)
	fmt.Println(payload)
	if payload[0] != 10 {
		t.Error("expected data byte 0 to be 10")
	}
	if payload[1] != 200 {
		t.Error("expected data byte 1 to be 200")
	}
}

func TestGetMessageById(t *testing.T) {
	data := ecu_model.MessageMap{}
	err := loaders.LoadDbc("tesla_can.dbc", &data)
	if err != nil {
		t.Error(err)
	}
	message := can_database.GetMessageById(data, 1160)
	if message == nil {
		t.Error("Returned message from GetMessageById was nil, should be DAS_steeringControl")
	}
	if message.Name != "DAS_steeringControl" {
		t.Error("Returned message from GetMessageById was not DAS_steeringControl")
	}

	// Try non-existant message
	badMessage := can_database.GetMessageById(data, 1)
	if badMessage != nil {
		t.Error("GetMessageById returned a message when it should returned nil")
	}
}
