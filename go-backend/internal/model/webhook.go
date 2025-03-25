package model

// WebhookPayload represents the top-level webhook payload
type WebhookPayload struct {
	Message WebhookMessage `json:"message"`
}

// WebhookMessage represents a webhook message from the API
type WebhookMessage struct {
	Timestamp                          int64               `json:"timestamp,omitempty"`
	Type                               string              `json:"type"`
	Call                               Call                `json:"call"`
	Customer                           Customer            `json:"customer"`
	Status                             string              `json:"status,omitempty"`
	EndedReason                        string              `json:"endedReason,omitempty"`
	InboundPhoneCallDebuggingArtifacts *DebuggingArtifacts `json:"inboundPhoneCallDebuggingArtifacts,omitempty"`
	PhoneNumber                        *PhoneNumber        `json:"phoneNumber,omitempty"`
}

// Call represents a call in the webhook message
type Call struct {
	ID                  string   `json:"id"`
	OrgID               string   `json:"orgId"`
	CreatedAt           string   `json:"createdAt,omitempty"`
	UpdatedAt           string   `json:"updatedAt,omitempty"`
	Type                string   `json:"type"`
	Status              string   `json:"status"`
	PhoneCallProvider   string   `json:"phoneCallProvider,omitempty"`
	PhoneCallProviderID string   `json:"phoneCallProviderId,omitempty"`
	PhoneCallTransport  string   `json:"phoneCallTransport,omitempty"`
	PhoneNumberID       string   `json:"phoneNumberId,omitempty"`
	AssistantID         *string  `json:"assistantId,omitempty"`
	SquadID             *string  `json:"squadId,omitempty"`
	Customer            Customer `json:"customer"`
}

// Customer represents a customer in the webhook message
type Customer struct {
	Number string `json:"number"`
}

// PhoneNumber represents a phone number configuration
type PhoneNumber struct {
	ID                                   string    `json:"id"`
	OrgID                                string    `json:"orgId"`
	AssistantID                          *string   `json:"assistantId,omitempty"`
	Number                               string    `json:"number"`
	CreatedAt                            string    `json:"createdAt,omitempty"`
	UpdatedAt                            string    `json:"updatedAt,omitempty"`
	StripeSubscriptionID                 *string   `json:"stripeSubscriptionId,omitempty"`
	TwilioAccountSid                     string    `json:"twilioAccountSid,omitempty"`
	TwilioAuthToken                      string    `json:"twilioAuthToken,omitempty"`
	StripeSubscriptionStatus             *string   `json:"stripeSubscriptionStatus,omitempty"`
	StripeSubscriptionCurrentPeriodStart *string   `json:"stripeSubscriptionCurrentPeriodStart,omitempty"`
	Name                                 string    `json:"name,omitempty"`
	CredentialID                         *string   `json:"credentialId,omitempty"`
	ServerURL                            *string   `json:"serverUrl,omitempty"`
	ServerURLSecret                      *string   `json:"serverUrlSecret,omitempty"`
	TwilioOutgoingCallerID               *string   `json:"twilioOutgoingCallerId,omitempty"`
	SipURI                               *string   `json:"sipUri,omitempty"`
	Provider                             string    `json:"provider,omitempty"`
	FallbackForwardingPhoneNumber        *string   `json:"fallbackForwardingPhoneNumber,omitempty"`
	FallbackDestination                  *string   `json:"fallbackDestination,omitempty"`
	SquadID                              *string   `json:"squadId,omitempty"`
	CredentialIDs                        *[]string `json:"credentialIds,omitempty"`
	NumberE164CheckEnabled               *bool     `json:"numberE164CheckEnabled,omitempty"`
	Authentication                       *string   `json:"authentication,omitempty"`
	Server                               *string   `json:"server,omitempty"`
	UseClusterSip                        *bool     `json:"useClusterSip,omitempty"`
	Status                               string    `json:"status,omitempty"`
	ProviderResourceID                   *string   `json:"providerResourceId,omitempty"`
}

// DebuggingArtifacts represents debugging information for failed calls
type DebuggingArtifacts struct {
	Error                    string                 `json:"error,omitempty"`
	AssistantRequestError    string                 `json:"assistantRequestError,omitempty"`
	AssistantRequestResponse map[string]interface{} `json:"assistantRequestResponse,omitempty"`
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
	CallTypeInboundPhoneCall CallType = "inboundPhoneCall"
)
