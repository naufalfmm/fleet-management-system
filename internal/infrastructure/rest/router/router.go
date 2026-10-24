package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/rest/controller/vehiclelocation"
	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/rest/middleware"
)

type Router struct {
	vehicleLocationController vehiclelocation.Controller

	middleware middleware.Middleware
}

func NewRouter(vehicleLocationController vehiclelocation.Controller, middleware middleware.Middleware) Router {
	return Router{
		vehicleLocationController: vehicleLocationController,
		middleware:                middleware,
	}
}

func (r Router) Register(ge *gin.Engine) {
	ge.RedirectFixedPath = true
	ge.Use(r.middleware.PanicRecover(), r.middleware.ImplementCors())

	auth := ge.Group("/vehicles")
	auth.GET("/:vehicle_id/location", r.vehicleLocationController.GetLatestLocation)
	auth.GET("/:vehicle_id/history", r.vehicleLocationController.GetLocationHistories)
}
