package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"local_library/config"
	"local_library/core/domain"
	"net/http"
)

type UserHandler struct {
	config *config.Config
}

func NewUserHandler(config *config.Config) *UserHandler {
	return &UserHandler{config: config}
}

func (h *UserHandler) Register(ctx *gin.Context) {

	requestUrl := fmt.Sprintf("http://%s:%s/user", h.config.CentralLibHost, h.config.CentralLibPort)
	// Read the raw request body
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	//Call to central lib
	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	// Read the response body
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare response body"})
		return
	}

	switch response.StatusCode {
	case 201:
		var newUser domain.User
		// Unmarshal the JSON data into the User struct
		err = json.Unmarshal(responseBody, &newUser)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		ctx.JSON(http.StatusCreated, newUser)
		return
	case 409:
		response := string(responseBody)
		ctx.JSON(http.StatusConflict, response[1:len(response)-1])
		return
	}

}
