package routestop

type GetRestForm struct {
	RouteID             int64
	Direction           string
	StopOrderBiggerThan int32
}
