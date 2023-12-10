package payloads

type RequestPostUsers struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=12,max=15"`
	FullName    string `json:"full_name" validate:"required,min=3,max=60"`
	Password    string `json:"password" validate:"required,min=6,max=64"`
}

type ResponsePostUsers struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
