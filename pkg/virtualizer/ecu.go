package virtualizer

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/can_database"
	"github.com/bit-dream/go-virtual/pkg/ecu_model"
)

// StartAllVirtualEcus starts all virtual ecus provided a list of virtual ecus
func StartAllVirtualEcus(ecus []ecu_model.VirtualEcu) (chan bool, error) {

	stop := make(chan bool)

	for _, ecu := range ecus {
		err := VirtualizeEcu(ecu, stop)
		if err != nil {
			stop <- true
			return nil, fmt.Errorf("")
		}
	}

	return stop, nil
}

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

		signals := can_database.MarshalMessageToVirtualSignals(dbcMessage, message)
		if message.Signals != nil {
			signals = can_database.OverrideVirtualSignals(signals, *message.Signals)
		}
		signalList := ecu_model.MarshalSignalMapToArray(signals)
		if signals == nil || len(signals) == 0 {
			return fmt.Errorf("an error occured while generating signal list to be sent by the virtual ecu")
		}

		rate := message.TransmissionOptions.TransmitRate
		if rate != nil {
			go PeriodicTimer(*rate, stop, func() {
				payload := can_database.MarshalSignalsToPayload(signalList)
				fmt.Println(payload)
			})
		} else {
			return fmt.Errorf("only periodic transmission is available currently, please include a TransmitRate")
		}
	}

	return nil
}
