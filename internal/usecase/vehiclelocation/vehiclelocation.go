package vehiclelocation

import (
	"context"

	"github.com/naufalfmm/fleet-management-system/config"
	"github.com/naufalfmm/fleet-management-system/internal/model/dao"
	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/common"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/messagequeue"
	"github.com/naufalfmm/fleet-management-system/internal/repository/routestop"
	"github.com/naufalfmm/fleet-management-system/internal/repository/vehicleassignment"
	"github.com/naufalfmm/fleet-management-system/internal/repository/vehiclelocation"
	"github.com/shopspring/decimal"
)

type Usecase interface {
	InsertLocationFromFleet(ctx context.Context, req dto.InsertLocationRequest) error
	GetLatestVehicleLocation(ctx context.Context, vehicleID string) (dao.VehicleLocation, error)
	GetVehicleLocationHistories(ctx context.Context, req dto.GetVehicleLocationHistoriesRequest) ([]dao.VehicleLocation, error)
}

type usecase struct {
	vehicleLocationRepository   vehiclelocation.Repository
	vehicleAssignmentRepository vehicleassignment.Repository
	routeStopRepository         routestop.Repository

	messageQueue messagequeue.MessageQueue

	conf config.Config
}

func NewUsecase(vehicleLocationRepository vehiclelocation.Repository, vehicleAssignmentRepository vehicleassignment.Repository, routeStopRepository routestop.Repository, messageQueue messagequeue.MessageQueue, conf config.Config) Usecase {
	return &usecase{
		vehicleLocationRepository:   vehicleLocationRepository,
		vehicleAssignmentRepository: vehicleAssignmentRepository,
		routeStopRepository:         routeStopRepository,

		messageQueue: messageQueue,

		conf: conf,
	}
}

func (u *usecase) InsertLocationFromFleet(ctx context.Context, req dto.InsertLocationRequest) error {
	vehicleLoc, err := u.vehicleLocationRepository.Create(ctx, dao.VehicleLocation{
		VehicleID: req.VehicleID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: req.Timestamp,
	})
	if err != nil {
		return err
	}

	if err := u.checkGeofencing(ctx, req.VehicleID, vehicleLoc); err != nil {
		return err
	}

	return nil
}

func (u *usecase) checkGeofencing(ctx context.Context, vehicleID string, currentVehicleLocation dao.VehicleLocation) error {
	vehicleAssignment, err := u.vehicleAssignmentRepository.GetByVehicleID(ctx, vehicleID)
	if err != nil {
		return nil
	}

	if vehicleAssignment.AssignedRouteID == nil {
		return nil
	}

	routeGetRestForm := routestop.GetRestForm{
		RouteID:   *vehicleAssignment.AssignedRouteID,
		Direction: *vehicleAssignment.AssignedDirection,
	}

	if vehicleAssignment.CurrentRouteStopID != nil {
		routeStop, err := u.routeStopRepository.GetByID(ctx, *vehicleAssignment.CurrentRouteStopID)
		if err != nil {
			return nil
		}

		routeGetRestForm.StopOrderBiggerThan = routeStop.StopOrder
	}

	routeStops, err := u.routeStopRepository.GetRest(ctx, routeGetRestForm)
	if err != nil {
		return nil
	}

	for _, routeStop := range routeStops {
		dist := common.CalculateDistanceFast(currentVehicleLocation.Latitude, currentVehicleLocation.Longitude, routeStop.Latitude, routeStop.Longitude)
		if dist.LessThanOrEqual(decimal.NewFromFloat(u.conf.GeofencingRadiusMeter)) {
			if _, err := u.vehicleAssignmentRepository.UpdateCurrentStop(ctx, vehicleAssignment.ID, routeStop.ID); err != nil {
				return nil
			}

			go func() {
				u.messageQueue.Publish(ctx, u.conf.RabbitMqGeofencingExchange, u.conf.RabbitMqGeofencingQueue, dto.GeofenceEventMessage{
					VehicleID: vehicleID,
					Event:     "geofence_entry",
					Location: dto.GeofenceEventMessageLocation{
						Latitude:  currentVehicleLocation.Latitude,
						Longitude: currentVehicleLocation.Longitude,
					},
					Timestamp: currentVehicleLocation.Timestamp,
				})
			}()
		}
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
