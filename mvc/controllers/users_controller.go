// + build ignore
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/prosline/micro-services/mvc/services"
	"github.com/prosline/micro-services/mvc/utils"
)

func GetUser(resp http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	log.Printf("Processing user_id %v", userId)

	user, apiErr := services.GetUser(userId)
	if err != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
