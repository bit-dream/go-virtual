package model_test

import (
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"testing"
)

func TestModelLoad(t *testing.T) {

	_, err := ecu_model.LoadEcu("test1.ecu")
	if err != nil {
		t.Error(err)
	}
}
