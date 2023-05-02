package virtualizer

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/can_database"
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"go.einride.tech/can/pkg/descriptor"
)

// VirtualizeEcu Virtualizes the provided ecu model. Each message within the virtual ecu will run as a separate go
// routine
func VirtualizeEcu(ecu ecu_model.VirtualEcu, stop chan bool) error {

	if ecu.Messages == nil {
		return fmt.Errorf("provided ecu contained messages that were nil")
	}

	for _, message := range *ecu.Messages {

		dbcData := *ecu.DbcData
		dbcMessage, ok := dbcData[message.Name]
		if !ok {
			return fmt.Errorf("could not find message %s in DBC data", message.Name)
		}

		signals := MarshalMessageToVirtualSignals(dbcMessage, message)
		if message.Signals != nil {
			signals = OverrideVirtualSignals(signals, *message.Signals)
		}
		signalList := getSignalMapValues(signals)
		if signals == nil || len(signals) == 0 {
			return fmt.Errorf("an error occured while generating signal list to be sent by the virtual ecu")
		}

		rate := message.TransmissionOptions.TransmitRate
		if rate != nil {
			go PeriodicTimer(*rate, stop, func() {
				payload := can_database.GeneratePayloadFromSignals(signalList)
				fmt.Println(payload)
			})
		} else {
			return fmt.Errorf("only periodic transmission is available currently, please include a TransmitRate")
		}
	}

	return nil
}

type SignalMap map[string]ecu_model.VirtualSignal

func getSignalMapValues(m SignalMap) []ecu_model.VirtualSignal {
	values := make([]ecu_model.VirtualSignal, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func MarshalMessageToVirtualSignals(message descriptor.Message, virtualMessage ecu_model.VirtualMessage) SignalMap {

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

func OverrideVirtualSignals(signals SignalMap, overrideSignals []ecu_model.VirtualSignal) SignalMap {
	for _, overrideSignal := range overrideSignals {
		virtualSignal, hasSignal := signals[overrideSignal.Name]
		if hasSignal {
			overrideSignal.SignalDefinition = virtualSignal.SignalDefinition
			signals[overrideSignal.Name] = overrideSignal
		}
	}
	return signals
}
