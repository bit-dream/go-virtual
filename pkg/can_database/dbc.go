package can_database

import (
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"go.einride.tech/can/pkg/descriptor"
)

func GetMessageById(dbc loaders.DbcData, id uint32) *descriptor.Message {

	for _, value := range dbc.Messages {
		if value.ID == id {
			return &value
		}
	}
	return nil
}
