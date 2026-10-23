package vehiclelocation

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/naufalfmm/fleet-management-system/internal/domain/vehiclelocation"
	"github.com/naufalfmm/fleet-management-system/internal/model/dao"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/common"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/txmanager"
	"github.com/shopspring/decimal"
)

type Repository interface {
	Create(ctx context.Context, vehicleLocation dao.VehicleLocation) (dao.VehicleLocation, error)
	GetLatest(ctx context.Context, vehicleID string) (dao.VehicleLocation, error)
	GetHistories(ctx context.Context, form GetHistoriesForm) ([]dao.VehicleLocation, error)
}

type repository struct {
	querier vehiclelocation.Querier
}

func NewRepository(querier vehiclelocation.Querier) Repository {
	return &repository{
		querier: querier,
	}
}

func newVehicleLocation(vehLoc vehiclelocation.VehicleLocation) (dao.VehicleLocation, error) {
	flLat, err := vehLoc.Latitude.Float64Value()
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	flLong, err := vehLoc.Longitude.Float64Value()
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	return dao.VehicleLocation{
		ID:        vehLoc.ID,
		VehicleID: vehLoc.VehicleID,
		Latitude:  decimal.NewFromFloat(flLat.Float64),
		Longitude: decimal.NewFromFloat(flLong.Float64),
		Timestamp: vehLoc.Timestamp,
		CreatedAt: vehLoc.CreatedAt.Time,
	}, nil
}

func (r *repository) Create(ctx context.Context, vehicleLocation dao.VehicleLocation) (dao.VehicleLocation, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = vehiclelocation.New(tx)
	}

	numLat, err := common.NewDecimalToPgTypeNumeric(vehicleLocation.Latitude)
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	numLong, err := common.NewDecimalToPgTypeNumeric(vehicleLocation.Longitude)
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	vl, err := q.Create(ctx, vehiclelocation.CreateParams{
		VehicleID: vehicleLocation.VehicleID,
		Latitude:  numLat,
		Longitude: numLong,
		Timestamp: vehicleLocation.Timestamp,
	})
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	return newVehicleLocation(vl)
}

func (r *repository) GetLatest(ctx context.Context, vehicleID string) (dao.VehicleLocation, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = vehiclelocation.New(tx)
	}

	vl, err := q.GetLatest(ctx, vehicleID)
	if err != nil {
		return dao.VehicleLocation{}, err
	}

	return newVehicleLocation(vl)
}

func (r *repository) GetHistories(ctx context.Context, form GetHistoriesForm) ([]dao.VehicleLocation, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = vehiclelocation.New(tx)
	}

	vls, err := q.GetHistory(ctx, vehiclelocation.GetHistoryParams{
		VehicleID:     form.VehicleID,
		TimestampFrom: form.TimestampFrom,
		TimestampTo:   form.TimestampTo,
	})
	if err != nil {
		return nil, err
	}

	vehicleLocs := make([]dao.VehicleLocation, len(vls))
	for i, vl := range vls {
		vehicleLoc, err := newVehicleLocation(vl)
		if err != nil {
			return nil, err
		}

		vehicleLocs[i] = vehicleLoc
	}

	return vehicleLocs, nil
}
