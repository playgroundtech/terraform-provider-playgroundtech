package client

type Application struct {
	ID        int       `json:"applicant_id"`
	PhoneNr   string	`json:"phone_number,omitempty"`
	Email     string    `json:"email,omitempty"`
	LinkedIn  string    `json:"linkedin,omitempty"`
	Github    string    `json:"github"`
	HomePage  string    `json:"homepage"`
}