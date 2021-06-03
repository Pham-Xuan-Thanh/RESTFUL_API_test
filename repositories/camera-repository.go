package repositories

import (
	"golang-example/entities"

	"gorm.io/gorm"
)

type CameraRepository interface {
	FindAll(boxID string) (*[]entities.Camera, error)
	FindOneById(id string, boxID string) (*entities.Camera, error)
	InsertCamera(createCamera entities.Camera) (*entities.Camera, error)
	UpdateCamera(updateCamera entities.Camera) (*entities.Camera, error)
	DeleteCamera(id string, boxID string) (*entities.Camera, error)
}

type camerarepo struct {
	db *gorm.DB
}

func (c camerarepo) FindAll(boxID string) (*[]entities.Camera, error) {
	var cameras *[]entities.Camera
	result := c.db.Where("box_id=?", boxID).Preload("Box").Find(&cameras)

	if result.Error != nil {
		return nil, result.Error
	}
	return cameras, nil
}

func (c camerarepo) FindOneById(id string, boxID string) (*entities.Camera, error) {
	var camera *entities.Camera
	result := c.db.Table("cameras").Where("box_id=?", boxID).Preload("Box").First(&camera, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return camera, nil
}

func (c camerarepo) InsertCamera(createCamera entities.Camera) (*entities.Camera, error) {
	result := c.db.Table("cameras").Create(&createCamera)
	c.db.Preload("Box").Find(&createCamera)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createCamera, nil
}

func (c camerarepo) UpdateCamera(updateCamera entities.Camera) (*entities.Camera, error) {
	result := c.db.Table("cameras").Updates(&updateCamera)
	c.db.Preload("Box").Find(&updateCamera)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateCamera, nil
}

func (c camerarepo) DeleteCamera(id string, boxID string) (*entities.Camera, error) {
	camera, err := c.FindOneById(id, boxID)
	if err != nil {
		return nil, err
	}

	result := c.db.Delete(&camera, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return camera, nil
}

func NewCameraRepository(db *gorm.DB) CameraRepository {
	return &camerarepo{db}
}
