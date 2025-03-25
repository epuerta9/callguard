package service

import "github.com/epuerta/callguard/go-backend/internal/model"

// MinimalWebhookPayload represents the minimal structure needed to identify the webhook type
type MinimalWebhookPayload struct {
	Type string `json:"type"`
}

// VAPIForwardPayload represents the incoming webhook payload
type VAPIForwardPayload struct {
	Timestamp   int64              `json:"timestamp"`
	Type        string             `json:"type"`
	Call        model.Call         `json:"call"`
	Customer    model.Customer     `json:"customer"`
	Status      string             `json:"status,omitempty"`
	EndedReason string             `json:"endedReason,omitempty"`
	PhoneNumber *model.PhoneNumber `json:"phoneNumber,omitempty"`
}

// VAPIForwardResponse represents the response payload
type VAPIForwardResponse struct {
	Destination ForwardDestination `json:"destination"`
}

// ForwardDestination represents the destination configuration
type ForwardDestination struct {
	Type                   string  `json:"type"`
	Message                string  `json:"message"`
	Number                 string  `json:"number"`
	NumberE164CheckEnabled bool    `json:"numberE164CheckEnabled"`
	CallerID               *string `json:"callerId,omitempty"`
	Extension              *string `json:"extension,omitempty"`
}
