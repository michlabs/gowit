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

// Intents is a helper method that returns intent fied of the meaning if have, 
// otherwise returns empty string
func (m *Meaning) Intent() string {
	var intent string
	if m.Entities["intent"] != nil {
		t := m.Entities["intent"][0]["value"]
		if str, ok := t.(string); ok {
			intent = str
		}
	}
	return intent
}