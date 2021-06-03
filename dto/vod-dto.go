package dto

import "time"

type CreateVodDto struct {
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	DeviceId  int       `gorm:"column:device_id" json:"device_id"`
	Name      string    `gorm:"column:name" json:"name"`
	State     string    `gorm:"column:state" json:"state"`
	Type      string    `gorm:"column:type" json:"type"`
	Size      int64     `gorm:"column:size" json:"size"`
	DashPath  string    `gorm:"column:dash_path" json:"dash_path"`
	StartTime time.Time `gorm:"DEFAULT:'current_timestamp'" json:"start_time"`
	EndTime   time.Time `gorm:"DEFAULT:'current_timestamp'" json:"end_time"`
}

type UpdateVodDto struct {
	Name     string    `gorm:"column:name" json:"name"`
	State    string    `gorm:"column:state" json:"state"`
	Size     int64     `gorm:"column:size" json:"size"`
	DashPath string    `gorm:"column:dash_path" json:"dash_path"`
	EndTime  time.Time `gorm:"DEFAULT:'current_timestamp'" json:"end_time"`
}

type FindVodDto struct {
	State string `gorm:"column:state" json:"state"`
}
