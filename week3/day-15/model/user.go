package model

type Article struct {
	ID      int
	Status  bool
	Title   string
	Content string
}

type Users struct{
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UsersResponse struct{
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}