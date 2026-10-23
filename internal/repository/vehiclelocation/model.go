package vehiclelocation

type GetHistoriesForm struct {
	VehicleID     string
	TimestampFrom int64
	TimestampTo   int64
}
