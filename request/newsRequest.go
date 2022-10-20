package request

type NewsRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FindParameter struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
