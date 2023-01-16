package userdto

type DeleteUserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateUserRequest struct {
	// Name     string `json:"name" form:"name"`
	// Password string `json:"password" form:"password"`
	// Phone    string `json:"phone" form:"phone"`
	Image string `json:"image" form:"image"`
}
