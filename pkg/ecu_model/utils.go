package ecu_model

// MarshalSignalMapToArray creates an array of virtual signals from a SignalMap object
func MarshalSignalMapToArray(m SignalMap) []VirtualSignal {
	values := make([]VirtualSignal, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
