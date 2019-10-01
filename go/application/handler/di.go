package handler

import (
	"go/domain/service"
	"go/infra"
)

var circleService service.CircleInterface

func init() {
	circleRepo := infra.NewCircleDB(infra.DB)
	circleService = service.NewCircleService(circleRepo)
}
