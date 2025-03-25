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
func (s *VapiService) CreateAssistant(ctx context.Context, systemPrompt string, endCallPhrases []string) (AssistantResponse, error) {
	// payload := &api.CreateAssistantDto{
	// 	FirstMessage: api.String("Hi there, I see you're calling about ms. Johnson. How can I help you today?"),
	// 	Model: &api.CreateAssistantDtoModel{
	// 		OpenAiModel: &api.OpenAiModel{
	// 			Model: "gpt-4o-mini",
	// 			Messages: []*api.OpenAiMessage{
	// 				{
	// 					Role:    api.OpenAiMessageRoleSystem,
	// 					Content: &systemPrompt,
	// 				},
	// 			},
	// 		},
	// 	},
	// 	Voice: &api.CreateAssistantDtoVoice{
	// 		ElevenLabsVoice: &api.ElevenLabsVoice{
	// 			VoiceId:         api.String("alloy"),
	// 			Stability:       api.Float64(0.5),
	// 			SimilarityBoost: api.Float64(0.5),
	// 		},
	// 	},
	// 	ServerMessages: []api.CreateAssistantDtoServerMessagesItem{
	// 		api.CreateAssistantDtoServerMessagesItem("function-call"),
	// 		api.CreateAssistantDtoServerMessagesItem("end-of-call-report"),
	// 	},
	// 	EndCallPhrases: endCallPhrases,
	// }

	jsonResp := AssistantResponse{
		Assistant: Assistant{
			Voice: Voice{
				VoiceID:                "3RMkGuPAdoghqmNaBi0l",
				Provider:               "11labs",
				Stability:              0.5,
				SimilarityBoost:        0.75,
				Model:                  "eleven_multilingual_v2",
				FillerInjectionEnabled: false,
			},
			Model: Model{
				Model:   "gpt-4o-mini",
				ToolIDs: []string{"e2a4fd36-4b4d-4152-9b31-588372bc626c"},
				Messages: []Message{
					{
						Role:    "system",
						Content: fmt.Sprintf("you are a helpful assistant"),
					},
				},
				Provider:                  "openai",
				EmotionRecognitionEnabled: false,
			},
			RecordingEnabled:       true,
			FirstMessage:           "Hi thanks for calling ms. Johnson, is there any message I can take for her?",
			VoicemailMessage:       "voicemailMessage",
			EndCallFunctionEnabled: true,
			EndCallMessage:         "endCallMessage",
			Transcriber: Transcriber{
				Model:    "nova-2",
				Language: "multi",
				Provider: "deepgram",
			},
			ServerMessages: []string{
				"transfer-destination-request",
				"end-of-call-report",
				"function-call",
			},
			EndCallPhrases: []string{
				"goodbye",
				"talk to you soon",
				"adios",
				"bye",
				"see you later",
				"hasta luego",
				"que tengas un buen dia",
				"fin de la llamada",
				"end of the call",
			},
			NumWordsToInterruptAssistant: 2,
			BackgroundSound:              "office",
			BackchannelingEnabled:        false,
			BackgroundDenoisingEnabled:   false,
			StartSpeakingPlan: StartSpeakingPlan{
				SmartEndpointingEnabled: true,
			},
		},
	}

	log.Printf("Created assistant with ID")
	return jsonResp, nil
}

// ListAssistants retrieves all available assistants
func (s *VapiService) ListAssistants(ctx context.Context) ([]*api.Assistant, error) {
	return s.client.Assistants.List(ctx, &api.AssistantsListRequest{
		Limit: api.Float64(10),
	})
}
