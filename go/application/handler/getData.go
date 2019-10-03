package handler

import (
	"go/application/response"
	"go/domain/model"
	"strings"

	"log"
	"net/http"
)

type getCircleRequest struct {
	Keyword string `json:"keyword"`
}

type getCircleResponse struct {
	Result []model.Circle `json:"result"`
}

func GetCircle(w http.ResponseWriter, r *http.Request) {
	keyvalue := strings.Split(r.URL.String(), "?")
	value := strings.Split(keyvalue[1], "=")
	print(value)
	circles, err := circleService.GetCircle(value[1])
	if err != nil {
		log.Println(err)
	}

	response.Success(w, getCircleResponse{Result: circles})
}
