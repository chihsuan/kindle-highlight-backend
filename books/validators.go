package books

type CreateValidator struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Author string `form:"author" json:"author" binding:"required"`
}
