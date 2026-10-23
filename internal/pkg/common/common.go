package common

import (
	"fmt"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

func NewDecimalToPgTypeNumeric(d decimal.Decimal) (pgtype.Numeric, error) {
	var num pgtype.Numeric

	err := num.Scan(d.String())
	if err != nil {
		return pgtype.Numeric{}, fmt.Errorf("failed to convert decimal to numeric: %w", err)
	}

	return num, nil
}

func CalculateDistanceFast(lat1, lon1, lat2, lon2 decimal.Decimal) decimal.Decimal {
	earthRadiusInMeters := decimal.NewFromFloat(6371000.0)

	pi := decimal.NewFromFloat(math.Pi)
	straightAngle := decimal.NewFromFloat(180.)

	lat1Rad := lat1.Mul(pi).Div(straightAngle)
	lat2Rad := lat2.Mul(pi).Div(straightAngle)

	dLat := lat2.Sub(lat1).Mul(pi).Div(straightAngle)
	dLon := lon2.Sub(lon1).Mul(pi).Div(straightAngle)

	meanLat := lat1Rad.Add(lat2Rad).Div(decimal.NewFromFloat(2.))

	x := dLon.Mul(meanLat.Cos())
	y := dLat

	sqrt, _ := x.Mul(x).Add(y.Mul(y)).Sqrt()

	return sqrt.Mul(earthRadiusInMeters)
}
