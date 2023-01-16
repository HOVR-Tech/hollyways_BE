package userdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"password"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	Role  string `json:"role"`
}

type DeleteUserResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
