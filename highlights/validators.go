package highlights

type CreateValidator struct {
	BookID     uint     `form:"bookID" json:"bookID" binding:"required"`
	UserID     uint     `form:"userID" json:"userID" binding:"required"`
	Highlights []string `form:"highlights" json:"highlights" binding:"required"`
}

type GetAllValidator struct {
	BookID uint `form:"bookID" json:"bookID" binding:"required"`
}
