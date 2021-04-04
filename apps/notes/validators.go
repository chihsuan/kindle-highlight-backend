package notes

type CreateValidator struct {
	UserID  uint   `form:"userID" json:"userID" binding:"required"`
	BookID  uint   `form:"bookID" json:"bookID" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type GetAllValidator struct {
	UserID uint `form:"userID" json:"userID" binding:"required"`
	BookID uint `form:"bookID" json:"bookID" binding:"required"`
}
