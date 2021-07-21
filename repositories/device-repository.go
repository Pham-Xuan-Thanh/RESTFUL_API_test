package repositories

import (
	"database/sql"
	"golang-example/entities"
	
	"gorm.io/gorm"
)

type DeviceRepository interface {
	FindAll(boxID string) (*[]entities.Device, error)
	FindOneById(id string, boxID string) (*entities.Device, error)
	InsertDevice(createDevice entities.Device) (*entities.Device, error)
	UpdateDevice(updateDevice entities.Device) (*entities.Device, error)
	DeleteDevice(id string, boxID string) (*entities.Device, error)

	GetAll() (*[]entities.Device, error)
	FindAllDeviceByFilter(service_id string) (*[]entities.Device, error)
}

type devicerepo struct {
	db *gorm.DB
}

func (c devicerepo) FindAll(boxID string) (*[]entities.Device, error) {
	var devices *[]entities.Device
	result := c.db.Where("box_id=?", boxID).Preload("Box").Find(&devices)

	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}

func (c devicerepo) FindOneById(id string, boxID string) (*entities.Device, error) {
	var device *entities.Device
	result := c.db.Table("devices").Where("box_id=?", boxID).First(&device, id)
	if result.Error != nil {
		return nil,  result.Error
	}
	return device, nil
}

func (c devicerepo) InsertDevice(createDevice entities.Device) (*entities.Device, error) {
	result := c.db.Table("devices").Create(&createDevice)
	c.db.Preload("Box").Find(&createDevice)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createDevice, nil
}

func (c devicerepo) UpdateDevice(updateDevice entities.Device) (*entities.Device, error) {
	result := c.db.Table("devices").Updates(&updateDevice)
	c.db.Preload("Box").Find(&updateDevice)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateDevice, nil
}

func (c devicerepo) DeleteDevice(id string, boxID string) (*entities.Device, error) {
	device, err := c.FindOneById(id, boxID)
	if err != nil {
		return nil, err 
	}

	result := c.db.Delete(&device, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return device, nil
}

func (c devicerepo) GetAll() (*[]entities.Device, error) {
	var devices *[]entities.Device
	result := c.db.Table("devices").Find(&devices)

	if result.Error != nil {
		return nil, result.Error
	}

	return devices, nil
}

func (c devicerepo) FindAllDeviceByFilter(service_id string) (*[]entities.Device, error) {
	var storage_files *[]entities.Device
	result := c.db.Raw("SELECT v.id, v.service_id, v.box_id, v.location_id, v.id, v.video_coding_standard, v.streaming_protocol, v.stream_name, v.stream_uri, v.stream_endpoint, v.is_recording_set, v.is_PTZ_enabled, v.is_connected, v.status, v.is_active, v.deleted_at, v.created_at, v.updated_at FROM devices v WHERE v.service_id = @service_id ORDER BY v.id ASC", sql.Named("service_id", service_id)).Find(&storage_files)
	if result.Error != nil {
		return nil, result.Error
	}
	return storage_files, nil
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &devicerepo{db}
}
