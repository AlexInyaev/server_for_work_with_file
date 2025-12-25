package formatter

type TextWriteRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Text  string `json:"text"`
}
type TextWriteResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Text  string `json:"text"`
}
