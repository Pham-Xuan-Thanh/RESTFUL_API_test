package dto

type CreateCameraPortalDto struct {
	Name string `json:"name" form:"name" binding:"required"`

	StreamName string `json:"stream_name" form:"stream_name" binding:"required"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"required,url"`

	ServiceID string `json:"service_id" form:"service_id" binding:"omitempty"`

	BoxID int64 `json:"box_id" form:"box_id" binding:"required"`
}

type CreateCameraDto struct {
	Name string `json:"name" form:"name" binding:"required"`

	StreamName string `json:"stream_name" form:"stream_name" binding:"required"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"required,url"`

	ServiceID string `json:"service_id" form:"service_id" binding:"required"`

	BoxID int64 `json:"box_id" form:"box_id" binding:"required"`
}

type UpdateCameraDto struct {
	ID int64 `json:"id" form:"id" binding:"required"`

	Name string `json:"name" form:"name"`

	StreamName string `json:"stream_name" form:"stream_name"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"omitempty,url"`

	IsActive int `json:"is_active" form:"is_active"`

	BoxID int64 `json:"box_id" form:"box_id" binding:"omitempty"`
}
