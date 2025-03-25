package service

type Result struct {
	ToolCallId string `json:"toolCallId"`
	Result     string `json:"result"`
}

type Response struct {
	Results []Result `json:"results"`
}

type ServerEvent struct {
	Message struct {
		Type string `json:"type"`
		Call Call   `json:"call"`
	} `json:"message"`
}

type ServerToolCall struct {
	Message struct {
		Type      string     `json:"type"`
		Call      Call       `json:"call"`
		ToolCalls []ToolCall `json:"toolCalls"`
	} `json:"message"`
}

type AssistantResponse struct {
	Assistant Assistant `json:"assistant"`
}

// Principal struct
type ToolCall struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name      string      `json:"name"`
	Arguments interface{} `json:"arguments"`
}

type ToolWithTool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
	Async    bool     `json:"async"`
	Server   Server   `json:"server"`
	// Messages []Message `json:"messages"`
	ToolCall ToolCall `json:"toolCall"`
}

type Server struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
}

type CustomerPhoneCall struct {
	Number string `json:"number"`
}
type Call struct {
	ID          string            `json:"id"`
	OrgID       string            `json:"orgId"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
	Type        string            `json:"type"`
	Status      string            `json:"status"`
	AssistantID string            `json:"assistantId"`
	WebCallUrl  string            `json:"webCallUrl,omitempty"`
	Customer    CustomerPhoneCall `json:"customer,omitempty"`
}

type Artifact struct {
	Messages                []ArtifactMessage    `json:"messages"`
	MessagesOpenAIFormatted []OpenAIFormattedMsg `json:"messagesOpenAIFormatted"`
}

type ArtifactMessage struct {
	Role             string     `json:"role"`
	Message          string     `json:"message"`
	Time             float64    `json:"time"`
	EndTime          float64    `json:"endTime"`
	SecondsFromStart float64    `json:"secondsFromStart"`
	Source           string     `json:"source"`
	Duration         float64    `json:"duration,omitempty"`
	ToolCalls        []ToolCall `json:"toolCalls,omitempty"`
}

type OpenAIFormattedMsg struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

type Customer struct {
	Number string `json:"number"`
}

// Assistant represents the "assistant" field
type Assistant struct {
	Voice                        Voice             `json:"voice"`
	Model                        Model             `json:"model"`
	RecordingEnabled             bool              `json:"recordingEnabled"`
	FirstMessage                 string            `json:"firstMessage"`
	VoicemailMessage             string            `json:"voicemailMessage"`
	EndCallFunctionEnabled       bool              `json:"endCallFunctionEnabled"`
	EndCallMessage               string            `json:"endCallMessage"`
	Transcriber                  Transcriber       `json:"transcriber"`
	ClientMessages               []string          `json:"clientMessages"`
	ServerMessages               []string          `json:"serverMessages"`
	EndCallPhrases               []string          `json:"endCallPhrases"`
	NumWordsToInterruptAssistant int               `json:"numWordsToInterruptAssistant"`
	BackgroundSound              string            `json:"backgroundSound"`
	BackchannelingEnabled        bool              `json:"backchannelingEnabled"`
	BackgroundDenoisingEnabled   bool              `json:"backgroundDenoisingEnabled"`
	StartSpeakingPlan            StartSpeakingPlan `json:"startSpeakingPlan"`
}

// Voice represents the "voice" field within "assistant"
type Voice struct {
	VoiceID                string  `json:"voiceId"`
	Provider               string  `json:"provider"`
	Stability              float64 `json:"stability"`
	SimilarityBoost        float64 `json:"similarityBoost"`
	Model                  string  `json:"model"`
	FillerInjectionEnabled bool    `json:"fillerInjectionEnabled"`
}

// Model represents the "model" field within "assistant"
type Model struct {
	Model                     string    `json:"model"`
	ToolIDs                   []string  `json:"toolIds"`
	Messages                  []Message `json:"messages"`
	Provider                  string    `json:"provider"`
	EmotionRecognitionEnabled bool      `json:"emotionRecognitionEnabled"`
}

// Message represents each entry in the "messages" array within "model"
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Transcriber represents the "transcriber" field within "assistant"
type Transcriber struct {
	Model    string `json:"model"`
	Language string `json:"language"`
	Provider string `json:"provider"`
}

type PhoneNumber struct {
	ID                                   string `json:"id"`
	OrgID                                string `json:"orgId"`
	AssistantID                          any    `json:"assistantId"`
	Number                               string `json:"number"`
	StripeSubscriptionID                 any    `json:"stripeSubscriptionId"`
	TwilioAccountSid                     string `json:"twilioAccountSid"`
	TwilioAuthToken                      string `json:"twilioAuthToken"`
	StripeSubscriptionStatus             any    `json:"stripeSubscriptionStatus"`
	StripeSubscriptionCurrentPeriodStart any    `json:"stripeSubscriptionCurrentPeriodStart"`
	Name                                 string `json:"name"`
	CredentialID                         any    `json:"credentialId"`
	ServerURL                            any    `json:"serverUrl"`
	ServerURLSecret                      any    `json:"serverUrlSecret"`
	TwilioOutgoingCallerID               any    `json:"twilioOutgoingCallerId"`
	SipURI                               any    `json:"sipUri"`
	Provider                             string `json:"provider"`
	FallbackForwardingPhoneNumber        any    `json:"fallbackForwardingPhoneNumber"`
	FallbackDestination                  any    `json:"fallbackDestination"`
	SquadID                              any    `json:"squadId"`
	CredentialIds                        any    `json:"credentialIds"`
	NumberE164CheckEnabled               any    `json:"numberE164CheckEnabled"`
	Authentication                       any    `json:"authentication"`
}

type AssistantRequest struct {
	Message struct {
		Type        string      `json:"type"`
		Call        Call        `json:"call"`
		PhoneNumber PhoneNumber `json:"phoneNumber"`
		Customer    Customer    `json:"customer"`
	} `json:"message"`
}

type EndCallReport struct {
	Message struct {
		Type            string      `json:"type"`
		Call            Call        `json:"call"`
		PhoneNumber     PhoneNumber `json:"phoneNumber"`
		Customer        Customer    `json:"customer"`
		DurationMs      int         `json:"durationMs"`
		DurationSeconds float64     `json:"durationSeconds"`
		DurationMinutes float64     `json:"durationMinutes"`
		Summary         string      `json:"summary"`
		Transcript      string      `json:"transcript"`
		EndedReason     string      `json:"endedReason"`
		Cost            float64     `json:"cost"`
	} `json:"message"`
}

type StartSpeakingPlan struct {
	SmartEndpointingEnabled bool `json:"smartEndpointingEnabled"`
}

type KnowledgeBase struct {
	TopK     int      `json:"topK"`
	FileIds  []string `json:"fileIds"`
	Provider string   `json:"provider"`
}
