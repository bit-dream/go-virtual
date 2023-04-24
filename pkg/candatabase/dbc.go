package candatabase

import "go.einride.tech/can/pkg/descriptor"

func GetMessageById(dbc DbcData, id uint32) *descriptor.Message {

	for _, value := range dbc.Messages {
		if value.ID == id {
			return &value
		}
	}
	return nil
}
