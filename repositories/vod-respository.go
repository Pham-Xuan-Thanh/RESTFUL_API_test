package repositories

/**
 * @author : Donald Duck Trieu
 * @created : 4/23/21, Friday
**/
import (
	"golang-example/entities"
	"database/sql"

	"gorm.io/gorm"
)

type VodRepository interface {
	AddRecordFile(createCamera entities.Vods) (*entities.Vods, error)
	UpdateRecordFile(uuid string, updateVods entities.Vods) (*entities.Vods, error)
	FindAllRecordFile() (*[]entities.Vods, error)
	FindAllVodByFilter(state string) (*[]entities.Vods, error)
}

type vodrepo struct {
	db *gorm.DB
}

func (c vodrepo) AddRecordFile(createVods entities.Vods) (*entities.Vods, error) {
	result := c.db.Table("vods").Create(&createVods)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createVods, nil
}

func (c vodrepo) UpdateRecordFile(uuid string, updateVods entities.Vods) (*entities.Vods, error) {
	result := c.db.Table("vods").Where("uuid = ?", uuid).Updates(&updateVods)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateVods, nil
}

func (c vodrepo) FindAllRecordFile() (*[]entities.Vods, error) {
	var vods *[]entities.Vods
	result := c.db.Table("vods").Find(&vods)

	if result.Error != nil {
		return nil, result.Error
	}
	return vods, nil
}

func (c vodrepo) FindAllVodByFilter(state string) (*[]entities.Vods, error) {
	var storage_files *[]entities.Vods
	result := c.db.Raw("SELECT v.uuid, v.state, v.size, v.dash_path, v.start_time, v.end_time, v.deleted_at, v.created_at, v.updated_at FROM vods v WHERE v.state = @state ORDER BY v.end_time ASC", sql.Named("state", state)).Find(&storage_files)
	if result.Error != nil {
		return nil, result.Error
	}
	return storage_files, nil
}

func NewVODRepository(db *gorm.DB) VodRepository {
	return &vodrepo{db}
}
