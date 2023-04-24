package virtualizer

type ExternalDataTransfer struct {
	Name          string
	Ecu           string
	Command       int
	CommandBuffer *[]int
}
