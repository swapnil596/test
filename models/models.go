package models

type API struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type Health struct {
	Status string `json:"status" example:"running"`
}

type APIStatus struct {
	Success	bool `json:"success" example:"true"`
}