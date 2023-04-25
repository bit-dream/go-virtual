package candatabase_test

import (
	"github.com/bit-dream/go-virtual/pkg/can_database"
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"testing"
)

func TestLoadDbcFile(t *testing.T) {
	_, err := loaders.LoadDbc("tesla_can.dbc")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadingOfMessages(t *testing.T) {
	data, err := loaders.LoadDbc("tesla_can.dbc")
	if err != nil {
		t.Error(err)
	}
	if len(data.Messages) == 0 {
		t.Error("no messages collected during parsing of dbc file")
	}
	if len(data.Messages) != 42 {
		t.Error("expected number of parsed messages from dbc file is incorrect, should have gotten 42")
	}
}

func TestGetMessageById(t *testing.T) {
	data, err := loaders.LoadDbc("tesla_can.dbc")
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
