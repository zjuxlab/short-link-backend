package controller

type loginReq struct {
	Name string `json:"name" validate:"required"`
}
