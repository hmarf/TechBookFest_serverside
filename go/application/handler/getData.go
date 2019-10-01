package handler

import (
	"encoding/json"
	"go/application/response"
	"go/domain/model"

	"io/ioutil"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		response.BadRequest(w, "Invalid Request Body")
		return
	}

	var requestBody getCircleRequest
	json.Unmarshal(body, &requestBody)

	if requestBody.Keyword == "" {
		response.BadRequest(w, "Invalid Request Body")
		return
	}

	circles, err := circleService.GetCircle(requestBody.Keyword)
	if err != nil {
		log.Println(err)
	}

	response.Success(w, getCircleResponse{Result: circles})
}
