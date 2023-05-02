package can_database

import (
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/descriptor"
)

// GetMessageById returns a message by CAN ID
func GetMessageById(dbc ecu_model.MessageMap, id uint32) *descriptor.Message {

	for _, value := range dbc {
		if value.ID == id {
			return &value
		}
	}
	return nil
}

// MarshalSignalsToPayload provided a list of virtual signals will provide the subsequent CAN data payload
func MarshalSignalsToPayload(signals []ecu_model.VirtualSignal) can.Data {
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

// MarshalMessageToVirtualSignals concat a message and virtual message's signals into a signal mapping, with the signal
// name being the keys to the map
func MarshalMessageToVirtualSignals(message descriptor.Message, virtualMessage ecu_model.VirtualMessage) ecu_model.SignalMap {

	var defaultValue int = 0
	if virtualMessage.DefaultValueForSignals != nil {
		defaultValue = *virtualMessage.DefaultValueForSignals
	}

	virtualSignals := make(map[string]ecu_model.VirtualSignal)
	for _, signal := range message.Signals {
		_, keyAlreadyExists := virtualSignals[signal.Name]
		if keyAlreadyExists {
			continue
		}
		newSignal := ecu_model.VirtualSignal{
			Name:             signal.Name,
			DefaultValue:     defaultValue,
			SignalDefinition: signal,
		}
		virtualSignals[signal.Name] = newSignal
	}

	return virtualSignals
}

// OverrideVirtualSignals overrides the provided signal map with override signals. This is typically use to overwrite
// the data provided by the raw dbc signals with the signals defined by the virtual ecu
func OverrideVirtualSignals(signals ecu_model.SignalMap, overrideSignals []ecu_model.VirtualSignal) ecu_model.SignalMap {
	for _, overrideSignal := range overrideSignals {
		virtualSignal, hasSignal := signals[overrideSignal.Name]
		if hasSignal {
			overrideSignal.SignalDefinition = virtualSignal.SignalDefinition
			signals[overrideSignal.Name] = overrideSignal
		}
	}
	return signals
}
