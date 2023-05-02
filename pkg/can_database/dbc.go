package can_database

import (
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/descriptor"
)

func GetMessageById(dbc ecu_model.MessageMap, id uint32) *descriptor.Message {

	for _, value := range dbc {
		if value.ID == id {
			return &value
		}
	}
	return nil
}
func GeneratePayloadFromSignals(signals []ecu_model.VirtualSignal) can.Data {
	data := can.Data{}
	if len(signals) == 0 {
		return data
	}
	for _, signal := range signals {
		var sigVal int64
		var usigVal uint64
		if signal.Value == nil {
			sigVal = (int64)(signal.DefaultValue)
			usigVal = (uint64)(signal.DefaultValue)
		} else {
			sigVal = (int64)(*signal.Value)
			usigVal = (uint64)(*signal.Value)
		}
		if signal.SignalDefinition.IsBigEndian {
			if signal.SignalDefinition.IsSigned {
				data.SetSignedBitsBigEndian(
					signal.SignalDefinition.Start,
					signal.SignalDefinition.Length,
					sigVal,
				)
			} else {
				data.SetUnsignedBitsBigEndian(
					signal.SignalDefinition.Start,
					signal.SignalDefinition.Length,
					usigVal,
				)
			}
		} else {
			if signal.SignalDefinition.IsSigned {
				data.SetSignedBitsLittleEndian(
					signal.SignalDefinition.Start,
					signal.SignalDefinition.Length,
					sigVal,
				)
			} else {
				data.SetUnsignedBitsLittleEndian(
					signal.SignalDefinition.Start,
					signal.SignalDefinition.Length,
					usigVal,
				)
			}
		}
	}

	return data
}
