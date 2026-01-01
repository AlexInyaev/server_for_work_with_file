package formatter

type TextWriteRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Text       string `json:"text"`
	Directory  string `json:"directory"`
	RecordType string `json:"recordType"`
}
type TextWriteResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Text       string `json:"text"`
	Directory  string `json:"directory"`
	RecordType string `json:"recordType"`
}
