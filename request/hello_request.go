package request

type HelloQueryRequest struct {
	Title   string `json:"title,omitempty" form:"title"`
	Content string `json:"content,omitempty" form:"content"`
}

type HelloBodyRequest struct {
	Title   string `json:"title" form:"title" validate:"required" example:"This title"`
	Content string `json:"content" form:"content" validate:"required" example:"This content"`
}

type HelloFormDataRequest struct {
	Title   string `json:"title" form:"title" validate:"required" example:"This title"`
	Content string `json:"content" form:"content" validate:"required" example:"This content"`
}
