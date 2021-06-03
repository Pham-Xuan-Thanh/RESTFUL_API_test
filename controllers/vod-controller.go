package controllers

/**
 * @author : Donald Trieu
 * @created : 4/23/21, Friday
**/
import (
	"golang-example/dto"
	"golang-example/helpers"
	"golang-example/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type VodController interface {
	GetAllRecordFile(ctx *gin.Context)
	AddRecordFile(ctx *gin.Context)
	UpdateRecordFile(ctx *gin.Context)
	GetAllVodByFilter(ctx *gin.Context)
}

type vodcontroller struct {
	VodService services.VODService
}

// @tags boxes
// @Summary Insert single vod
// @Description get all vod
// @Accept  json
// @Produce  json
// @Param device_id query int false "string"
// @success 200 {object} helpers.Response{data=[]entities.Vods}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/box/vod [get]
func (c *vodcontroller) AddRecordFile(ctx *gin.Context) {
	var VodCreateDto dto.CreateVodDto
	errDTO := ctx.ShouldBind(&VodCreateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.VodService.AddRecordFile(VodCreateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}

// @tags boxes
// @Summary Insert single vod
// @Description get all vod
// @Accept  json
// @Produce  json
// @Param device_id query int false "string"
// @success 200 {object} helpers.Response{data=[]entities.Vods}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/box/vod [get]
func (c *vodcontroller) GetAllRecordFile(ctx *gin.Context) {
	result, err := c.VodService.GetAllFileRecord()
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		response := helpers.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	}
}

// @tags boxes
// @Summary Insert single vod
// @Description get all vod
// @Accept  json
// @Produce  json
// @Param device_id query int false "string"
// @success 200 {object} helpers.Response{data=[]entities.Vods}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/box/vod [get]
func (c *vodcontroller) UpdateRecordFile(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	var UpdateVodDto dto.UpdateVodDto
	ctx.ShouldBind(&UpdateVodDto)
	if len(strings.TrimSpace(uuid)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query uuid not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.VodService.UpdateFileRecord(uuid, UpdateVodDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// @tags boxes
// @Summary get all vod
// @Description get all vod
// @Accept  json
// @Produce  json
// @Param device_id query int false "int"
// @success 200 {object} helpers.Response{data=[]entities.Vods}
// @failure 400,404 {object} helpers.Response{}
// @Router /api/box/vod [get]
func (c *vodcontroller) GetAllVodByFilter(ctx *gin.Context) {
	state := ctx.Query("state")
	if len(strings.TrimSpace(state)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query playback not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.VodService.GetAllVodByFilter(state)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func NewVODController(VodService services.VODService) VodController {
	return &vodcontroller{
		VodService: VodService,
	}
}
