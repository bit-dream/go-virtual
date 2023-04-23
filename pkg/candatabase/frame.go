package candatabase

import (
	"fmt"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/descriptor"
)

func DecodeFrame(frame can.Frame, message descriptor.Message) {

	data := frame.Data
	signals := message.Signals

	for _, signal := range signals {
		if signal.IsBigEndian {
			if signal.IsSigned {
				bitfield := data.SignedBitsBigEndian(signal.Start, signal.Length)
				fmt.Println(bitfield)
			} else {
				bitfield := data.UnsignedBitsBigEndian(signal.Start, signal.Length)
				fmt.Println(bitfield)
			}
		} else {
			if signal.IsSigned {
				bitfield := data.SignedBitsLittleEndian(signal.Start, signal.Length)
				fmt.Println(bitfield)
			} else {
				bitfield := data.UnsignedBitsLittleEndian(signal.Start, signal.Length)
				fmt.Println(bitfield)
			}
		}
	}

}

func EncodeFrame(message descriptor.Message) {
	payload := can.Data{}
	fmt.Println(payload)
}
