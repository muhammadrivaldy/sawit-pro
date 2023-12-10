package payloads

type RequestPostUsersLogin struct {
	PhoneNumber string `json:"phone_number" validate:"required,phone-number,min=12,max=15"`
	Password    string `json:"password" validate:"required,password,min=6,max=64"`
}

type ResponsePostUsersLogin struct {
	Token string `json:"token"`
}
