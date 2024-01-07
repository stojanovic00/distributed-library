package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"local_library/config"
	"local_library/core/domain"
	"local_library/core/repo"
	"net/http"
)

type IssuingHandler struct {
	config      *config.Config
	issuingRepo repo.IssuingRepo
}

func NewIssuingHandler(config *config.Config, issuingRepo repo.IssuingRepo) *IssuingHandler {
	return &IssuingHandler{config: config, issuingRepo: issuingRepo}
}
func (h *IssuingHandler) GetAllIssuingRecords(ctx *gin.Context) {

	records, err := h.issuingRepo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, records)
}

func (h *IssuingHandler) RecordBookIssue(ctx *gin.Context) {
	var record domain.IssuingRecord

	err := ctx.ShouldBindJSON(&record)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if h.issuingRepo.CheckIfTaken(record.ISBN) {
		ctx.JSON(http.StatusBadRequest, "book already taken")
		return
	}

	requestUrl := fmt.Sprintf("http://%s:%s/user/%d/record-issue/", h.config.CentralLibHost, h.config.CentralLibPort, record.UserId)
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
	case 200:
		err := h.issuingRepo.RecordBookIssue(&record)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, string(err.Error()))
			return
		}
		ctx.Status(http.StatusOK)
		return
	case 409:
		response := string(responseBody)
		ctx.JSON(http.StatusConflict, response[1:len(response)-1])
		return
	case 404:
		response := string(responseBody)
		ctx.JSON(http.StatusNotFound, response[1:len(response)-1])
		return
	case 400:
		response := string(responseBody)
		ctx.JSON(http.StatusBadRequest, response[1:len(response)-1])
		return
	}

}
func (h *IssuingHandler) RecordBookReturn(ctx *gin.Context) {
	isbn := ctx.Param("isbn")

	userId, err := h.issuingRepo.RecordBookReturn(isbn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	requestUrl := fmt.Sprintf("http://%s:%s/user/%d/record-return/", h.config.CentralLibHost, h.config.CentralLibPort, userId)
	//Call to central lib
	response, err := http.Post(requestUrl, "application/json", nil)
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
	case 200:
		ctx.Status(http.StatusOK)
		return
	case 404:
		response := string(responseBody)
		ctx.JSON(http.StatusConflict, response[1:len(response)-1])
		return
	}
}
