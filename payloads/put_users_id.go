package payloads

type RequestPutUsersId struct {
	PhoneNumber string `json:"phone_number" validate:"required,phone-number,min=12,max=15"`
	FullName    string `json:"full_name" validate:"required,min=3,max=60"`
}
