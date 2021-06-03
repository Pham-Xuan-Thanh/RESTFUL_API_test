package services

import (
	"golang-example/dto"
	"golang-example/entities"
	"golang-example/repositories"

	"github.com/mashingan/smapping"
)

type CameraService interface {
	GetSingleCamera(id string, boxID string) (*entities.Camera, error)
	GetAllCamera(boxID string) (*[]entities.Camera, error)
	InsertCameraPortal(cameraCreateDto dto.CreateCameraPortalDto) (*entities.Camera, error)
	InsertCamera(cameraCreateDto dto.CreateCameraDto) (*entities.Camera, error)
	UpdateCamera(cameraUpdateDto dto.UpdateCameraDto) (*entities.Camera, error)
	DeleteCamera(id string, boxID string) (*entities.Camera, error)
}

type cameraservice struct {
	CameraReposi repositories.CameraRepository
}

func (c *cameraservice) InsertCameraPortal(cameraCreateDto dto.CreateCameraPortalDto) (*entities.Camera, error) {
	camera := entities.Camera{}
	err := smapping.FillStruct(&camera, smapping.MapFields(&cameraCreateDto))
	if err != nil {
		return nil, err
	}
	return c.CameraReposi.InsertCamera(camera)
}

func (c *cameraservice) DeleteCamera(id string, boxID string) (*entities.Camera, error) {
	return c.CameraReposi.DeleteCamera(id, boxID)
}

func (c *cameraservice) GetSingleCamera(id string, boxID string) (*entities.Camera, error) {
	return c.CameraReposi.FindOneById(id, boxID)
}

func (c *cameraservice) GetAllCamera(boxID string) (*[]entities.Camera, error) {
	return c.CameraReposi.FindAll(boxID)
}

func (c *cameraservice) UpdateCamera(cameraUpdateDto dto.UpdateCameraDto) (*entities.Camera, error) {
	camera := entities.Camera{}
	err := smapping.FillStruct(&camera, smapping.MapFields(&cameraUpdateDto))
	if err != nil {
		return nil, err
	}
	return c.CameraReposi.UpdateCamera(camera)
}

func (c *cameraservice) InsertCamera(cameraCreateDto dto.CreateCameraDto) (*entities.Camera, error) {
	camera := entities.Camera{}
	err := smapping.FillStruct(&camera, smapping.MapFields(&cameraCreateDto))
	if err != nil {
		return nil, err
	}
	return c.CameraReposi.InsertCamera(camera)
}

func NewCameraService(cameraReposi repositories.CameraRepository) CameraService {
	return &cameraservice{
		CameraReposi: cameraReposi,
	}
}
