package domain

type SenderRequest struct {
	To           string                 `json:"to"`
	TemplateName string                 `json:"template_name"`
	Variables    map[string]interface{} `json:"variables"`
}
