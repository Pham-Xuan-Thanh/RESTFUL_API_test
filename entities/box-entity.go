package entities

type Box struct {
	ID int64 `json:"id" form:"id" binding:"required"`

	TenantID int64 `json:"tenant_id" form:"tenant_id"`

	LocationID int64 `json:"location_id" form:"location_id"`

	BoxName string `json:"box_name" form:"box_name"`

	BoxIp string `json:"box_ip" form:"box_ip" binding:"omitempty,ip"`
}
