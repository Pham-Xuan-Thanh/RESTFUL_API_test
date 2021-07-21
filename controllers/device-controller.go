package controllers

import (
	"golang-example/dto"
	"golang-example/helpers"
	"golang-example/services"
	"net/http"

	// "strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeviceController interface {
	GetAllDevice(ctx *gin.Context)
	GetDeviceSingle(ctx *gin.Context)
	InsertDevice(ctx *gin.Context)
	UpdateDevice(ctx *gin.Context)
	DeleteDevice(ctx *gin.Context)

	GetAll(ctx *gin.Context)
	GetAllDeviceByFilter(ctx *gin.Context)
}

type devicecontroller struct {
	DeviceService services.DeviceService
}

func (d *devicecontroller) UpdateDevice(ctx *gin.Context) {
	var deviceUpdateDto dto.UpdateDeviceDto
	errDTO := ctx.ShouldBind(&deviceUpdateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := d.DeviceService.UpdateDevice(deviceUpdateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (d *devicecontroller) DeleteDevice(ctx *gin.Context) {
	id := ctx.Query("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", strings.TrimSpace(boxID) + "Param query box_id not found" +strings.TrimSpace(id) , helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		if len(strings.TrimSpace(id)) == 0 {
			res := helpers.BuildErrorResponse("Failed to process request", "", helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			result, err := d.DeviceService.DeleteDevice(id, boxID)
			if err != nil {
				res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
				ctx.JSON(http.StatusBadRequest, res)
			} else {
				response := helpers.BuildResponse(true, "OK", result)
				ctx.JSON(http.StatusOK, response)
			}
		}
	}
}

func (d *devicecontroller) InsertDevice(ctx *gin.Context) {
	var deviceCreateDto dto.CreateDeviceDto
	errDTO := ctx.ShouldBind(&deviceCreateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := d.DeviceService.InsertDevice(deviceCreateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}

func (d *devicecontroller) GetAllDevice(ctx *gin.Context) {
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := d.DeviceService.GetAllDevice(boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (d *devicecontroller) GetAll(ctx *gin.Context) {
	result, err := d.DeviceService.GetAll()
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		response := helpers.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func (d *devicecontroller) GetAllDeviceByFilter(ctx *gin.Context) {
	service_id := ctx.Query("service_id")
	if len(strings.TrimSpace(service_id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query playback not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := d.DeviceService.GetAllDeviceByFilter(service_id)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (d *devicecontroller) GetDeviceSingle(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id or id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := d.DeviceService.GetSingleDevice(id, boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func NewDeviceController(deviceService services.DeviceService) DeviceController {
	return &devicecontroller{
		DeviceService: deviceService,
	}
}
