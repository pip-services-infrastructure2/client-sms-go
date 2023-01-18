package version1

type SmsRecipientV1 struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Language string `json:"language"`
}
