package models

type StorageRequest struct {
	Key    string `json:"key" binding:"required" swaggertype:"string" example:"hello"`
	Value  string `json:"Value" binding:"required" swaggertype:"string" example:"world"`
	Status string `json:"status" swaggertype:"string" example:"SUCCESS" swaggerignore:"true"`
}

type GetRequest struct {
	Key   string `json:"key" binding:"required" swaggertype:"string" example:"hello"`
	Value string `json:"value" swaggertype:"string" example:"world"`
}
