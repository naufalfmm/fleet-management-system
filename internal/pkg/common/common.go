package common

import (
	"fmt"

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
