package vehicleassignment

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naufalfmm/fleet-management-system/internal/domain/vehicleassignment"
	"github.com/naufalfmm/fleet-management-system/internal/model/dao"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/txmanager"
)

type Repository interface {
	GetByVehicleID(ctx context.Context, vehicleID string) (dao.VehicleAssignment, error)
	UpdateCurrentStop(ctx context.Context, id int64, currentStopID int64) (dao.VehicleAssignment, error)
}

type repository struct {
	querier vehicleassignment.Querier
}

func NewRepository(querier vehicleassignment.Querier) Repository {
	return &repository{
		querier: querier,
	}
}

func newVehicleAssignment(va vehicleassignment.VehicleAssignment) (dao.VehicleAssignment, error) {
	vehAssignment := dao.VehicleAssignment{
		ID:        va.ID,
		VehicleID: va.VehicleID,
		CreatedAt: va.CreatedAt.Time,
	}

	if va.AssignedRouteID.Valid {
		vehAssignment.AssignedRouteID = &va.AssignedRouteID.Int64
	}

	if va.CurrentRouteStopID.Valid {
		vehAssignment.CurrentRouteStopID = &va.CurrentRouteStopID.Int64
	}

	if va.AssignedDirection.Valid {
		vehAssignment.AssignedDirection = &va.AssignedDirection.String
	}

	return vehAssignment, nil
}

func (r *repository) GetByVehicleID(ctx context.Context, vehicleID string) (dao.VehicleAssignment, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = vehicleassignment.New(tx)
	}

	va, err := q.GetByVehicleID(ctx, vehicleID)
	if err != nil {
		return dao.VehicleAssignment{}, err
	}

	return newVehicleAssignment(va)
}

func (r *repository) UpdateCurrentStop(ctx context.Context, id int64, currentStopID int64) (dao.VehicleAssignment, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = vehicleassignment.New(tx)
	}

	va, err := q.UpdateCurrentStop(ctx, vehicleassignment.UpdateCurrentStopParams{
		CurrentRouteStopID: pgtype.Int8{
			Int64: currentStopID,
			Valid: true,
		},
		ID: id,
	})
	if err != nil {
		return dao.VehicleAssignment{}, err
	}

	return newVehicleAssignment(va)
}
