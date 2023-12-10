package payloads

type ResponseGetUsersId struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
