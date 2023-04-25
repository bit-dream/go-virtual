package model_test

import (
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"testing"
)

func TestModelLoad(t *testing.T) {

	_, err := loaders.LoadEcu("test1.ecu")
	if err != nil {
		t.Error(err)
	}
}
