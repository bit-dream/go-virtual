package main

type VirtualModel struct {
	Messages []VirtualMessage `json:"messages"`
	Channels []Channel        `json:"channels"`
}

type Channel struct {
	Name           string          `json:"name"`
	BaudRate       int             `json:"baudRate"`
	ChannelOptions []ChannelOption `json:"channelOptions,omitempty"`
}

type VirtualMessage struct {
	Name                   string               `json:"name"`
	Dbc                    string               `json:"dbc"`
	Channel                string               `json:"channel"`
	DefaultValueForSignals *int                 `json:"defaultValueForSignals,omitempty"`
	Signals                *[]VirtualSignal     `json:"signals,omitempty"`
	TransmissionOptions    *TransmissionOptions `json:"transmissionOptions,omitempty"`
}

type VirtualSignal struct {
	Name                string               `json:"name"`
	DefaultValue        int                  `json:"defaultValue"`
	ValueQueue          *[]int               `json:"valueQueue,omitempty"`
	TransmissionOptions *TransmissionOptions `json:"transmissionOptions,omitempty"`
}

type TransmissionOptions struct {
	TransmitRate            *int      `json:"transmitRate,omitempty"`
	TriggeredByMessageIds   *[]int    `json:"triggeredByMessageIds,omitempty"`
	TriggeredByMessageNames *[]string `json:"triggeredByMessageNames,omitempty"`
	TriggeredBySignalNames  *[]string `json:"triggeredBySignalNames"`
}

type ChannelOption struct {
	Channel       string `json:"channel"`
	PacketSpacing int    `json:"packetSpacing"`
}

func LoadVirtualModel(vm VirtualModel) {

}
