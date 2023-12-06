package request

type UserCreateRequest struct {
	Name 		string `json:"name" validate:"required"`
	Email 		string `json:"email" validate:"required"`
	Adress 		string `json:"adress"`
	Phone 		string `json:"phone"`
}

type UserUpdateRequest struct {
	Name 		string `json:"name"`
	Adress 		string `json:"adress"`
	Phone 		string `json:"phone"`
}

type UserEmailRequest struct {
	Email 		string `json:"email" validate:"required"`
}