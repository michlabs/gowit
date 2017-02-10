package gowit

// WitError represents the error message returned by Wit.AI
type WitError struct {
	Error string `json:"error"`
	Code int `json:"code"`
}

// Context represents the context of conversation
type Context struct {
	ReferenceTime string `json:"reference_time"`
	Timezone      string `json:"timezone"`
}

// Meaning represents the extracted meaning from a sentence
type Meaning struct {
	MessageID string `json:"msg_id"`
	Text string `json:"_text"`
	Entities map[string][](map[string]interface{}) `json:"entities"`
}