package controllers

import (
	"golang-example/dto"
	"golang-example/helpers"
	"golang-example/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CameraController interface {
	GetAllCamera(ctx *gin.Context)
	GetCameraSingle(ctx *gin.Context)
	InsertCamera(ctx *gin.Context)
	UpdateCamera(ctx *gin.Context)
	DeleteCamera(ctx *gin.Context)

	GetAllCameraPortal(ctx *gin.Context)
	GetCameraSinglePortal(ctx *gin.Context)
	InsertCameraPortal(ctx *gin.Context)
	UpdateCameraPortal(ctx *gin.Context)
	DeleteCameraPortal(ctx *gin.Context)
}

type cameracontroller struct {
	CameraService services.CameraService
}

// @tags portal
// @Summary get all camera
// @Description get cameras
// @Accept  json
// @Produce  json
// @Param box_id query int false "int"
// @success 200 {object} helpers.Response{data=[]entities.Camera}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/portal/cameras [get]
func (c *cameracontroller) GetAllCameraPortal(ctx *gin.Context) {
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.GetAllCamera(boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// @tags portal
// @Summary get single camera from a box
// @Description get single camera from a box
// @Accept  json
// @Produce  json
// @Param id path int true "Camera ID"
// @Param box_id query int true "Box ID"
// @success 200 {object} helpers.Response{data=entities.Camera}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/portal/cameras/{id} [get]
func (c *cameracontroller) GetCameraSinglePortal(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id or id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.GetSingleCamera(id, boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// @tags portal
// @Summary add single camera
// @Description add single camera
// @Accept  json
// @Produce  json
// @Param request body dto.CreateCameraPortalDto true "CameraDto"
// @success 200 {object} helpers.Response{data=entities.Camera}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/portal/cameras [post]
func (c *cameracontroller) InsertCameraPortal(ctx *gin.Context) {
	var cameraCreateDto dto.CreateCameraPortalDto
	errDTO := ctx.ShouldBind(&cameraCreateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.InsertCameraPortal(cameraCreateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}

// @tags portal
// @Summary update single camera
// @Description update single camera
// @Accept  json
// @Produce  json
// @Param id path int true "Camera ID"
// @Param box_id query int true "Box ID"
// @Param request body dto.UpdateCameraDto true "updateCameraDto"
// @success 200 {object} helpers.Response{data=entities.Camera}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/portal/cameras/{id} [put]
func (c *cameracontroller) UpdateCameraPortal(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		var cameraUpdateDto dto.UpdateCameraDto
		errDTO := ctx.ShouldBind(&cameraUpdateDto)
		i64, _ := strconv.ParseInt(boxID, 10, 32)
		cameraUpdateDto.BoxID = i64
		if errDTO != nil {
			res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			result, err := c.CameraService.UpdateCamera(cameraUpdateDto)
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

// @tags portal
// @Summary delete single camera
// @Description delete single camera
// @Accept  json
// @Produce  json
// @Param id path int true "Camera ID"
// @Param box_id query int true "Box ID"
// @success 200 {object} helpers.Response{data=entities.Camera}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/portal/cameras/{id} [delete]
func (c *cameracontroller) DeleteCameraPortal(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		if len(strings.TrimSpace(id)) == 0 {
			res := helpers.BuildErrorResponse("Failed to process request", "", helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			result, err := c.CameraService.DeleteCamera(id, boxID)
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

func (c *cameracontroller) UpdateCamera(ctx *gin.Context) {
	var cameraUpdateDto dto.UpdateCameraDto
	errDTO := ctx.ShouldBind(&cameraUpdateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.UpdateCamera(cameraUpdateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c *cameracontroller) DeleteCamera(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		if len(strings.TrimSpace(id)) == 0 {
			res := helpers.BuildErrorResponse("Failed to process request", "", helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			result, err := c.CameraService.DeleteCamera(id, boxID)
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

func (c *cameracontroller) InsertCamera(ctx *gin.Context) {
	var cameraCreateDto dto.CreateCameraDto
	errDTO := ctx.ShouldBind(&cameraCreateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.InsertCamera(cameraCreateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}

func (c *cameracontroller) GetAllCamera(ctx *gin.Context) {
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.GetAllCamera(boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c *cameracontroller) GetCameraSingle(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id or id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.CameraService.GetSingleCamera(id, boxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func NewCameraController(cameraService services.CameraService) CameraController {
	return &cameracontroller{
		CameraService: cameraService,
	}
}
