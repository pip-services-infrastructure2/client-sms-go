package version1

type SmsMessageV1 struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
