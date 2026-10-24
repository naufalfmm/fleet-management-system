package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	PanicRecover() gin.HandlerFunc
	ImplementCors() gin.HandlerFunc
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}
