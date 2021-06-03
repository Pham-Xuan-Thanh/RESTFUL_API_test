package entities

import (
	"time"
)

/**
 * @author : Donald Duck Trieu
 * @created : 4/23/21, Friday
**/

type Vods struct {
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	DeviceId  int       `gorm:"column:device_id" json:"device_id"`
	Name      string    `gorm:"column:name" json:"name"`
	State     string    `gorm:"column:state" json:"state"`
	Type      string    `gorm:"column:type" json:"type"`
	Size      int64     `gorm:"column:size" json:"size"`
	DashPath  string    `gorm:"column:dash_path" json:"dash_path"`
	StartTime time.Time `gorm:"DEFAULT:'current_timestamp'" json:"start_time"`
	EndTime   time.Time `gorm:"DEFAULT:'current_timestamp'" json:"end_time"`
	DeletedAt time.Time `gorm:"DEFAULT:'current_timestamp'" json:"deleted_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
