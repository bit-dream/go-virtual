package virtualizer

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/can_database"
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
	"go.einride.tech/can/pkg/descriptor"
)

func VirtualizeEcu(ecu ecu_model.VirtualEcu, stop chan bool) error {

	//transmitter := socketcan.NewTransmitter(conn)
	for _, message := range *ecu.Messages {
		rate := message.TransmissionOptions.TransmitRate

		dbcData := *ecu.DbcData
		dbcMessage, ok := dbcData[message.Name]
		if !ok {
			return fmt.Errorf("could not find message %s in DBC data", message.Name)
		}

		signals := ConvertDbcMessageToVirtualSignals(dbcMessage, message)
		/*
			if message.Signals != nil {
				OverrideVirtualSignals(signals, *message.Signals)
			}
		*/
		signalList := getAllValues(signals)
		if signals == nil || len(signals) == 0 {
			return fmt.Errorf("an error occured while generating signal list to be sent by the virtual ecu")
		}

		if rate != nil {
			go PeriodicTimer(*rate, stop, func() {
				if signals == nil {
					fmt.Println("Error encountered were signals were nil: ", message)
					return
				}
				payload := can_database.GeneratePayloadFromSignals(signalList)
				fmt.Println(payload)
			})
		} else {
			return fmt.Errorf("only periodic transmission is available currently, please include a TransmitRate")
		}
	}

	return nil
}

func getAllValues(m map[string]ecu_model.VirtualSignal) []ecu_model.VirtualSignal {
	values := make([]ecu_model.VirtualSignal, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func ConvertDbcMessageToVirtualSignals(message descriptor.Message, virtualMessage ecu_model.VirtualMessage) map[string]ecu_model.VirtualSignal {

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

func OverrideVirtualSignals(signals map[string]ecu_model.VirtualSignal, overrideSignals []ecu_model.VirtualSignal) {
	for _, overrideSignal := range overrideSignals {
		_, hasSignal := signals[overrideSignal.Name]
		if hasSignal {
			signals[overrideSignal.Name] = overrideSignal
		}
	}
}

// GenerateCompleteSignalList will consolidate signals that are in the dbc file and signals defined by the Virtual ECU
// so there are no duplicates
func GenerateCompleteSignalList(dbcMessage descriptor.Message, virtualMessage ecu_model.VirtualMessage) []ecu_model.VirtualSignal {

	dbcSignals := dbcMessage.Signals
	updatedSignalList := make([]ecu_model.VirtualSignal, 0)

	virtualSignals := make([]ecu_model.VirtualSignal, 0)
	if virtualMessage.Signals != nil {
		virtualSignals = *virtualMessage.Signals
		for idx, _ := range virtualSignals {
			updatedSignalList = append(updatedSignalList, virtualSignals[idx])
		}
	} else {

	}

	for _, dbcSignal := range dbcSignals {
		if !signalIsInList(dbcSignal.Name, virtualSignals) {
			var defaultValue int = 0
			if virtualMessage.DefaultValueForSignals != nil {
				defaultValue = *virtualMessage.DefaultValueForSignals
			}

			newSignal := ecu_model.VirtualSignal{
				Name:             dbcSignal.Name,
				DefaultValue:     defaultValue,
				SignalDefinition: dbcSignal,
			}

			updatedSignalList = append(updatedSignalList, newSignal)
		}
	}

	return updatedSignalList
}

// signalIsInList returns true if the signal name is in the provided signal list
func signalIsInList(signalName string, signalList []ecu_model.VirtualSignal) bool {
	for _, signal := range signalList {
		if signalName == signal.Name {
			return true
		}
	}
	return false
}
