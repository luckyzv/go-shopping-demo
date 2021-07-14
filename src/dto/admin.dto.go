package dto

type AdminDto struct {
  PageSize int
  PageNum int

  Status string `json:"status" binding:"required"`
}
