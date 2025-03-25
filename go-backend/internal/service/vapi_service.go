package service

import (
	"context"
	"fmt"
	"log"

	api "github.com/VapiAI/server-sdk-go"
	vapiclient "github.com/VapiAI/server-sdk-go/client"
)

// VapiService handles Vapi API interactions
type VapiService struct {
	client *vapiclient.Client
}

// NewVapiService creates a new VapiService
func NewVapiService(client *vapiclient.Client) *VapiService {
	return &VapiService{
		client: client,
	}
}

// CreateAssistant creates a new Vapi assistant with the given configuration
func (s *VapiService) CreateAssistant(ctx context.Context, systemPrompt string, endCallPhrases []string) (*api.Assistant, error) {
	assistant, err := s.client.Assistants.Create(ctx, &api.CreateAssistantDto{
		FirstMessage: api.String("Hi there, I see you're calling about ms. Johnson. How can I help you today?"),
		Model: &api.CreateAssistantDtoModel{
			OpenAiModel: &api.OpenAiModel{
				Model: "gpt-4o-mini",
				Messages: []*api.OpenAiMessage{
					{
						Role:    api.OpenAiMessageRoleSystem,
						Content: &systemPrompt,
					},
				},
			},
		},
		Voice: &api.CreateAssistantDtoVoice{
			ElevenLabsVoice: &api.ElevenLabsVoice{
				VoiceId: &api.ElevenLabsVoiceId{
					ElevenLabsVoiceIdEnum: api.ElevenLabsVoiceIdEnumDrew,
				},
				Stability:       api.Float64(0.5),
				SimilarityBoost: api.Float64(0.5),
			},
		},
		ServerMessages: []api.CreateAssistantDtoServerMessagesItem{
			api.CreateAssistantDtoServerMessagesItem("function-call"),
			api.CreateAssistantDtoServerMessagesItem("end-of-call-report"),
		},
		EndCallPhrases: endCallPhrases,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create assistant: %w", err)
	}

	log.Printf("Created assistant with ID: %s", assistant.Id)
	return assistant, nil
}

// ListAssistants retrieves all available assistants
func (s *VapiService) ListAssistants(ctx context.Context) ([]*api.Assistant, error) {
	return s.client.Assistants.List(ctx, &api.AssistantsListRequest{
		Limit: api.Float64(10),
	})
}
