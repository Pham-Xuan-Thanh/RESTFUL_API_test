package services

import (
	"golang-example/dto"
	"golang-example/entities"
	"golang-example/repositories"

	"github.com/mashingan/smapping"
)

type VODService interface {
	AddRecordFile(CreateVodDto dto.CreateVodDto) (*entities.Vods, error)
	GetAllFileRecord() (*[]entities.Vods, error)
	UpdateFileRecord(uuid string, updateVodDto dto.UpdateVodDto) (*entities.Vods, error)
	GetAllVodByFilter(state string) (*[]entities.Vods, error)
}

type VODservice struct {
	VodReposi repositories.VodRepository
}

func (c *VODservice) AddRecordFile(CreateVodDto dto.CreateVodDto) (*entities.Vods, error) {
	vod := entities.Vods{}
	err := smapping.FillStruct(&vod, smapping.MapFields(&CreateVodDto))
	if err != nil {
		return nil, err
	}
	return c.VodReposi.AddRecordFile(vod)
}

func (c *VODservice) GetAllFileRecord() (*[]entities.Vods, error) {
	return c.VodReposi.FindAllRecordFile()
}

func (c *VODservice) UpdateFileRecord(uuid string, UpdateVodDto dto.UpdateVodDto) (*entities.Vods, error) {
	vod := entities.Vods{}
	err := smapping.FillStruct(&vod, smapping.MapFields(&UpdateVodDto))
	if err != nil {
		return nil, err
	}
	return c.VodReposi.UpdateRecordFile(uuid, vod)
}

func (c *VODservice) GetAllVodByFilter(state string) (*[]entities.Vods, error) {
	return c.VodReposi.FindAllVodByFilter(state)
}

func NewVODService(VodReposi repositories.VodRepository) VODService {
	return &VODservice{
		VodReposi: VodReposi,
	}
}
