package dao

import "time"

type Route struct {
	ID        int64
	Code      string
	Name      string
	CreatedAt time.Time
}
