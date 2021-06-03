package entities

import (
	"time"
)

type Camera struct {
	ID int64 `gorm:"primary_key:auto_increment" json:"id"`

	ServiceID string `gorm:"type:varchar(255);not null" json:"service_id"`

	BoxID int64 `gorm:"not null" json:"box_id"`

	LocationID int64 `gorm:"not null" json:"location_id"`

	Name string `gorm:"type:varchar(255);not null" json:"name"`

	VideoCodingStandard string `gorm:"type:char(10);default:'h.264'" json:"video_coding_standard"`

	StreamingProtocol string `gorm:"type:char(10);default:'rtsp'" json:"streaming_protocol"`

	StreamName string `gorm:"type:varchar(255);not null" json:"stream_name"`

	StreamUri string `gorm:"type:varchar(255);not null" json:"stream_uri"`

	StreamEndpoint string `gorm:"type:varchar(255);not null" json:"stream_endpoint"`

	IsRecordingSet int `gorm:"default:0" json:"is_recording_set"`

	IsPTZEnabled int `gorm:"default:0" json:"is_PTZ_enabled"`

	IsConnected int `gorm:"default:0" json:"is_connected"`

	Status string `gorm:"default:'disconnected'" json:"status"`

	IsActive int `gorm:"default:0" json:"is_active"`

	DeletedAt time.Time `gorm:"index" json:"deleted_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
