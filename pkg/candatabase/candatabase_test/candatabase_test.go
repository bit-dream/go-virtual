package candatabase_test

import (
	"github.com/bit-dream/go-virtual/pkg/candatabase"
	"testing"
)

func TestLoadDbcFile(t *testing.T) {
	_, err := candatabase.LoadDbc("tesla_can.dbc")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadingOfMessages(t *testing.T) {
	data, err := candatabase.LoadDbc("tesla_can.dbc")
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
