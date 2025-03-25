package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	vapiclient "github.com/VapiAI/server-sdk-go/client"
	"github.com/VapiAI/server-sdk-go/option"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestVapiService_CreateAssistant(t *testing.T) {
	// Load .env file
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Get API key from .env
	apiKey := os.Getenv("VAPI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping test: VAPI_API_KEY not found in .env file")
	}

	// Create real Vapi client
	client := vapiclient.NewClient(option.WithToken(apiKey))
	service := NewVapiService(client)

	tests := []struct {
		name           string
		systemPrompt   string
		endCallPhrases []string
		expectedError  bool
	}{
		{
			name:           "successful assistant creation",
			systemPrompt:   "You are a helpful assistant that can speak english and spanish.",
			endCallPhrases: []string{"goodbye", "adios"},
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assistant, err := service.CreateAssistant(context.Background(), tt.systemPrompt, tt.endCallPhrases)

			if tt.expectedError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, assistant)
			assert.NotEmpty(t, assistant.Id)

			// Clean up: Delete the assistant after test
			if assistant != nil {
				_, err := client.Assistants.Delete(context.Background(), assistant.Id)
				if err != nil {
					t.Logf("Warning: Failed to delete assistant %s: %v", assistant.Id, err)
				}
			}
		})
	}
}

func TestVapiService_ListAssistants(t *testing.T) {
	// Load .env file
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("VAPI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping test: VAPI_API_KEY not found in .env file")
	}

	client := vapiclient.NewClient(option.WithToken(apiKey))
	service := NewVapiService(client)

	assistants, err := service.ListAssistants(context.Background())
	fmt.Println(assistants)
	assert.NoError(t, err)
	// assert.NotNil(t, assistants)
}
