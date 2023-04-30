package model_test

import (
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"github.com/bit-dream/go-virtual/pkg/virtualizer"
	"testing"
)

func TestModelLoad(t *testing.T) {

	_, err := loaders.LoadEcu("test1.ecu")
	if err != nil {
		t.Error(err)
	}
}

func TestCompleteSignalList(t *testing.T) {
	ecu, err := loaders.LoadEcu("test1.ecu")
	if err != nil {
		t.Error(err)
	}
	dbcData := *ecu.DbcData
	virtualMessage := *ecu.Messages
	accelerationMsg := dbcData["AccelerationSensor"]
	newSignals := virtualizer.GenerateCompleteSignalList(accelerationMsg, virtualMessage[0])

	validSignals := [7]string{
		"SpprtVrblTrnsRpttnRtFrAcclrtnSns",
		"VrtclAcclrtnExRngeFigureOfMerit",
		"LngtdnlAcclrtnExRngFgureOfMerit",
		"LtrlAcclrtnExRangeFigureOfMerit",
		"VerticalAccelerationExRange",
		"LateralAccelerationExRange",
		"LongitudinalAccelerationExRange",
	}
	for _, signal := range newSignals {
		hasValidSignal := false
		for _, validSignal := range validSignals {
			if signal.Name == validSignal {
				hasValidSignal = true
			}
		}
		if !hasValidSignal {
			t.Error("malformed signal list created")
		}
	}
}
