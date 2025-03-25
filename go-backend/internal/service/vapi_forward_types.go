package service

// VAPIForwardPayload represents the incoming webhook payload
type VAPIForwardPayload struct {
	Type        string           `json:"type"`
	Artifact    ForwardArtifact  `json:"artifact"`
	Assistant   ForwardAssistant `json:"assistant"`
	PhoneNumber string           `json:"phoneNumber"`
	Customer    ForwardCustomer  `json:"customer"`
	Call        ForwardCall      `json:"call"`
}

// ForwardArtifact contains the conversation data
type ForwardArtifact struct {
	Messages                []ForwardMessage `json:"messages"`
	Transcript              string           `json:"transcript"`
	MessagesOpenAIFormatted []ForwardMessage `json:"messagesOpenAIFormatted"`
}

// ForwardMessage represents a single message in the conversation
type ForwardMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ForwardAssistant represents the assistant information
type ForwardAssistant struct {
	ID string `json:"id"`
}

// ForwardCustomer represents the customer information
type ForwardCustomer struct {
	ID string `json:"id"`
}

// ForwardCall represents the call information
type ForwardCall struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// VAPIForwardResponse represents the response payload
type VAPIForwardResponse struct {
	Destination ForwardDestination `json:"destination"`
}

// ForwardDestination represents the destination configuration
type ForwardDestination struct {
	Type                   string `json:"type"`
	Message                string `json:"message"`
	Number                 string `json:"number"`
	NumberE164CheckEnabled bool   `json:"numberE164CheckEnabled"`
	CallerID               string `json:"callerId"`
	Extension              string `json:"extension"`
}
