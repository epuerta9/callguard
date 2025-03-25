package model

// WebhookPayload represents the top-level webhook payload
type WebhookPayload struct {
	Message WebhookMessage `json:"message"`
}

// WebhookMessage represents a webhook message from the API
type WebhookMessage struct {
	Type                               string              `json:"type"`
	Call                               Call                `json:"call"`
	Customer                           Customer            `json:"customer"`
	Status                             string              `json:"status,omitempty"`
	EndedReason                        string              `json:"ended_reason,omitempty"`
	InboundPhoneCallDebuggingArtifacts *DebuggingArtifacts `json:"inbound_phone_call_debugging_artifacts,omitempty"`
}

// Call represents a call in the webhook message
type Call struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	OrgID  string `json:"org_id"`
}

// Customer represents a customer in the webhook message
type Customer struct {
	Number string `json:"number"`
}

// DebuggingArtifacts represents debugging information for failed calls
type DebuggingArtifacts struct {
	Error                 string `json:"error"`
	AssistantRequestError string `json:"assistant_request_error,omitempty"`
}

// CallStatus represents the status of a call
type CallStatus string

const (
	CallStatusRinging CallStatus = "ringing"
	CallStatusEnded   CallStatus = "ended"
)

// CallType represents the type of a call
type CallType string

const (
	CallTypeInboundPhoneCall CallType = "inbound_phone_call"
)
