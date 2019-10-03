package handler

import (
	"go/application/response"
	"go/domain/model"

	"log"
	"net/http"
)

type getallCircleResponse struct {
	Result []model.Circle `json:"result"`
}

func GetALLCircle(w http.ResponseWriter, r *http.Request) {

	circles, err := circleService.GetAllCircle()
	if err != nil {
		log.Println(err)
	}

	response.Success(w, getallCircleResponse{Result: circles})
}
