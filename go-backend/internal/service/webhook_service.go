package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/epuerta/callguard/go-backend/internal/model"
)

// WebhookService handles webhook processing
type WebhookService struct {
	callLogService *CallLogService
}

// NewWebhookService creates a new WebhookService
func NewWebhookService(callLogService *CallLogService) *WebhookService {
	return &WebhookService{
		callLogService: callLogService,
	}
}

// HandleWebhook processes incoming webhook messages based on their type
func (s *WebhookService) HandleWebhook(msg model.WebhookMessage) error {
	fmt.Printf("Processing webhook message of type: %s\n", msg.Type)

	switch msg.Type {
	case "assistant-request":
		return s.handleAssistantRequest(msg)
	case "status-update":
		return s.handleStatusUpdate(msg)
	default:
		return fmt.Errorf("unsupported webhook type: %s", msg.Type)
	}
}

// handleAssistantRequest processes assistant-request type messages
func (s *WebhookService) handleAssistantRequest(msg model.WebhookMessage) error {
	// Validate the call status
	fmt.Println("Call status:", msg.Call.Status)
	if msg.Call.Status != string(model.CallStatusRinging) {
		return errors.New("invalid call status for assistant request")
	}

	// Check if the caller number is spam
	isSpam := CheckSpam(msg.Customer.Number)
	if isSpam {
		log.Printf("Spam detected for number %s", msg.Customer.Number)
		return errors.New("spam number detected")
	}
	// Validate the call type
	if msg.Call.Type != string(model.CallTypeInboundPhoneCall) {
		return errors.New("invalid call type for assistant request")
	}

	// Create a new call log entry
	callLog := &model.CreateCallLogRequest{
		CallerNumber: msg.Customer.Number,
		CallDuration: 0, // Will be updated when call ends
	}

	// Save the call log
	_, err := s.callLogService.Create(nil, callLog, msg.Call.OrgID)
	if err != nil {
		return fmt.Errorf("failed to create call log: %w", err)
	}

	log.Printf("Processing assistant request for call %s from %s",
		msg.Call.ID,
		msg.Customer.Number,
	)

	return nil
}

// handleStatusUpdate processes status-update type messages
func (s *WebhookService) handleStatusUpdate(msg model.WebhookMessage) error {
	if msg.Status == string(model.CallStatusEnded) {
		log.Printf("Call %s ended with reason: %s", msg.Call.ID, msg.EndedReason)

		if msg.InboundPhoneCallDebuggingArtifacts != nil {
			log.Printf("Debug info - Error: %s", msg.InboundPhoneCallDebuggingArtifacts.Error)
			if msg.InboundPhoneCallDebuggingArtifacts.AssistantRequestError != "" {
				log.Printf("Assistant request error: %s", msg.InboundPhoneCallDebuggingArtifacts.AssistantRequestError)
			}
		}
	}

	return nil
}
