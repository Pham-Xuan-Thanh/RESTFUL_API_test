package services

import (
	"golang-example/dto"
	"golang-example/entities"
	"golang-example/repositories"

	"github.com/mashingan/smapping"
)

type DeviceService interface {
	GetSingleDevice(id string, boxID string) (*entities.Device, error)
	GetAllDevice(boxID string) (*[]entities.Device, error)
	InsertDevice(deviceCreateDto dto.CreateDeviceDto) (*entities.Device, error)
	UpdateDevice(deviceUpdateDto dto.UpdateDeviceDto) (*entities.Device, error)
	DeleteDevice(id string, boxID string) (*entities.Device, error)

	GetAll() (*[]entities.Device, error)
	GetAllDeviceByFilter(service_id string) (*[]entities.Device, error)
}

type deviceservice struct {
	DeviceReposi repositories.DeviceRepository
}

func (d *deviceservice) InsertDevice(deviceCreateDto dto.CreateDeviceDto) (*entities.Device, error) {
	device := entities.Device{}
	err := smapping.FillStruct(&device, smapping.MapFields(&deviceCreateDto))
	if err != nil {
		return nil, err
	}
	return d.DeviceReposi.InsertDevice(device)
}

func (d *deviceservice) DeleteDevice(id string, boxID string) (*entities.Device, error) {
	return d.DeviceReposi.DeleteDevice(id, boxID)
}

func (d *deviceservice) GetSingleDevice(id string, boxID string) (*entities.Device, error) {
	return d.DeviceReposi.FindOneById(id, boxID)
}

func (d *deviceservice) GetAllDevice(boxID string) (*[]entities.Device, error) {
	return d.DeviceReposi.FindAll(boxID)
}

func (d *deviceservice) UpdateDevice(deviceUpdateDto dto.UpdateDeviceDto) (*entities.Device, error) {
	device := entities.Device{}
	err := smapping.FillStruct(&device, smapping.MapFields(&deviceUpdateDto))
	if err != nil {
		return nil, err
	}
	return d.DeviceReposi.UpdateDevice(device)
}

// func (d *deviceservice) InsertDevice(deviceCreateDto dto.CreateDeviceDto) (*entities.Device, error) {
// 	device := entities.Device{}
// 	err := smapping.FillStruct(&device, smapping.MapFields(&deviceCreateDto))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return d.DeviceReposi.InsertDevice(device)
// }

func (d *deviceservice) GetAll() (*[]entities.Device, error) {
	return d.DeviceReposi.GetAll()
}

func (d *deviceservice) GetAllDeviceByFilter(service_id string) (*[]entities.Device, error) {
	return d.DeviceReposi.FindAllDeviceByFilter(service_id)
}

func NewDeviceService(deviceReposi repositories.DeviceRepository) DeviceService {
	return &deviceservice{
		DeviceReposi: deviceReposi,
	}
}
