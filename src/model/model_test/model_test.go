package model_test

import (
	"github.com/bit-dream/go-virtual/src/model"
	"testing"
)

var GlobalTestVm, _ = model.LoadModel("test.network")

func TestModelLoad(t *testing.T) {

	_, err := model.LoadModel("test.network")
	if err != nil {
		t.Error(err)
	}
}

func TestBadFileExtension(t *testing.T) {

	_, err := model.LoadModel("test.json")
	if err == nil {
		t.Error("LoadModel function did not return an error for inappropriate file type extension")
	}
}

func TestGetMessages(t *testing.T) {
	messages := model.GetMessages(*GlobalTestVm)
	if messages == nil {
		t.Error("Model returned no messages")
	}
}

func TestGroupMessagesByChannel(t *testing.T) {
	mapping := model.GetMessagesByChannels(*GlobalTestVm)
	if mapping == nil {
		t.Error("Model returned nil channels")
	}
}
