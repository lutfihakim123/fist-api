package book

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price string `json:"price" binding:"required,number"`
}