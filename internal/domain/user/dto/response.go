package dto

type UserResponse struct {
	ID        uint   `json:"id"`
	Token     string `json:"token,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LoginResponse struct {
	User UserResponse `json:"user"`
}
