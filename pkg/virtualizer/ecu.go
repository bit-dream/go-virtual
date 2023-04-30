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
		dbcSignals, ok := dbcData[message.Name]
		if !ok {
			return fmt.Errorf("could not find message %s in DBC data", message.Name)
		}

		signals := GenerateCompleteSignalList(dbcSignals, message)
		if signals == nil || len(signals) == 0 {
			return fmt.Errorf("signals returned nil in VirtualizeEcu")
		}

		if rate != nil {
			go PeriodicTimer(*rate, stop, func() {
				if signals == nil {
					fmt.Println("Error encountered were signals were nil: ", message)
					return
				}
				payload := can_database.GeneratePayloadFromSignals(signals)
				fmt.Println(payload)
			})
		} else {
			return fmt.Errorf("only periodic transmission is available currently, please include a TransmitRate")
		}
	}

	return nil
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
