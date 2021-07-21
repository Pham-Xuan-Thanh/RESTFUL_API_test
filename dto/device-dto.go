package dto

type CreateDeviceDto struct {
	Name string `json:"name" form:"name" binding:"required"`

	StreamName string `json:"stream_name" form:"stream_name" binding:"required"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"required,url"`

	FfmpegServiceID string `json:"ffmpeg_service_id" form:"ffmpeg_service_id" binding:"required"`

	PlaybackServiceID string `json:"playback_service_id" form:"playback_service_id" binding:"required"`

	FfmpegGlobalServiceID string `json:"ffmpeg_global_service_id" form:"ffmpeg_global_service_id" binding:"required"`

	BoxID int64 `json:"box_id" form:"box_id" binding:"required"`
}

type UpdateDeviceDto struct {
	ID int64 `json:"id" form:"id" binding:"required"`

	Name string `json:"name" form:"name"`

	StreamName string `json:"stream_name" form:"stream_name"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"omitempty,url"`

	IsActive int `json:"is_active" form:"is_active"`

	BoxID int64 `json:"box_id" form:"box_id" binding:"omitempty"`
}
