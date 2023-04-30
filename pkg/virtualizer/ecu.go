package virtualizer

/*
func VirtualizeEcu(ecu ecu_model.VirtualEcu, conn net.Conn, stop chan bool) error {

	transmitter := socketcan.NewTransmitter(conn)

	for _, message := range *ecu.Messages {

		rate := message.TransmissionOptions.TransmitRate
		if rate != nil {
			PeriodicTimer(*rate, stop, func() {
				signals := message.Signals
				data := can.Data{}
				if signals != nil {
					for _, signal := range *signals {
						if signal.Value != nil {

						} else {

						}
					}
				}
			})
		} else {
			return fmt.Errorf("only periodic transmission is available currently, please include a TransmitRate")
		}
	}

	return nil
}
*/
