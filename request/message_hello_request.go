package request

type MessageHelloRequest struct {
	Content string `form:"content" binding:"required" validate:"required" example:"Hello Encang Cutbray"`
}
