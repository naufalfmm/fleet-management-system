package routestop

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/naufalfmm/fleet-management-system/internal/domain/routestop"
	"github.com/naufalfmm/fleet-management-system/internal/model/dao"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/txmanager"
	"github.com/shopspring/decimal"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (dao.RouteStop, error)
	GetRest(ctx context.Context, form GetRestForm) ([]dao.RouteStop, error)
}

type repository struct {
	querier routestop.Querier
}

func NewRepository(querier routestop.Querier) Repository {
	return &repository{
		querier: querier,
	}
}

func newRouteStop(rs routestop.RouteStop) (dao.RouteStop, error) {
	flLat, err := rs.Latitude.Float64Value()
	if err != nil {
		return dao.RouteStop{}, err
	}

	flLon, err := rs.Longitude.Float64Value()
	if err != nil {
		return dao.RouteStop{}, err
	}

	return dao.RouteStop{
		ID:        rs.ID,
		RouteID:   rs.RouteID,
		Direction: rs.Direction,
		Name:      rs.Name,
		Latitude:  decimal.NewFromFloat(flLat.Float64),
		Longitude: decimal.NewFromFloat(flLon.Float64),
		StopOrder: rs.StopOrder,
		CreatedAt: rs.CreatedAt.Time,
	}, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (dao.RouteStop, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = routestop.New(tx)
	}

	rs, err := q.GetByID(ctx, id)
	if err != nil {
		return dao.RouteStop{}, err
	}

	return newRouteStop(rs)
}

func (r *repository) GetRest(ctx context.Context, form GetRestForm) ([]dao.RouteStop, error) {
	q := r.querier

	tx, ok := txmanager.ExtractTx[pgx.Tx](ctx)
	if ok {
		q = routestop.New(tx)
	}

	rss, err := q.GetRest(ctx, routestop.GetRestParams{
		RouteID:   form.RouteID,
		Direction: form.Direction,
		StopOrder: int32(form.StopOrderBiggerThan),
	})
	if err != nil {
		return nil, err
	}

	routeStops := make([]dao.RouteStop, len(rss))
	for i, rs := range rss {
		routeStop, err := newRouteStop(rs)
		if err != nil {
			return nil, err
		}

		routeStops[i] = routeStop
	}

	return routeStops, nil
}
