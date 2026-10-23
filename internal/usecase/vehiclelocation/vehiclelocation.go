package vehiclelocation

import (
	"context"

	"github.com/naufalfmm/fleet-management-system/internal/model/dao"
	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	"github.com/naufalfmm/fleet-management-system/internal/repository/vehiclelocation"
)

type Usecase interface {
	InsertLocationFromFleet(ctx context.Context, req dto.InsertLocationRequest) error
	GetLatestVehicleLocation(ctx context.Context, vehicleID string) (dao.VehicleLocation, error)
	GetVehicleLocationHistories(ctx context.Context, req dto.GetVehicleLocationHistoriesRequest) ([]dao.VehicleLocation, error)
}

type usecase struct {
	vehicleLocationRepository vehiclelocation.Repository
}

func NewUsecase(vehicleLocationRepository vehiclelocation.Repository) Usecase {
	return &usecase{
		vehicleLocationRepository: vehicleLocationRepository,
	}
}

func (u *usecase) InsertLocationFromFleet(ctx context.Context, req dto.InsertLocationRequest) error {
	_, err := u.vehicleLocationRepository.Create(ctx, dao.VehicleLocation{
		VehicleID: req.VehicleID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: req.Timestamp,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) GetLatestVehicleLocation(ctx context.Context, vehicleID string) (dao.VehicleLocation, error) {
	return u.vehicleLocationRepository.GetLatest(ctx, vehicleID)
}

func (u *usecase) GetVehicleLocationHistories(ctx context.Context, req dto.GetVehicleLocationHistoriesRequest) ([]dao.VehicleLocation, error) {
	return u.vehicleLocationRepository.GetHistories(ctx, vehiclelocation.GetHistoriesForm{
		VehicleID:     req.VehicleID,
		TimestampFrom: req.TimestampFrom,
		TimestampTo:   req.TimestampTo,
	})
}
