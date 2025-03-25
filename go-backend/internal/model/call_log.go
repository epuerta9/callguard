package model

// CallLog represents a call log record
type CallLog struct {
	ID             string  `json:"id"`
	UserID         string  `json:"user_id"`
	CallerNumber   string  `json:"caller_number"`
	CallDuration   int     `json:"call_duration"`
	RecordingURL   string  `json:"recording_url,omitempty"`
	Transcript     string  `json:"transcript,omitempty"`
	SentimentScore float64 `json:"sentiment_score,omitempty"`
	RiskScore      float64 `json:"risk_score,omitempty"`
	Flagged        bool    `json:"flagged"`
	Notes          string  `json:"notes,omitempty"`
	Tags           []Tag   `json:"tags,omitempty"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// CreateCallLogRequest represents a request to create a call log
type CreateCallLogRequest struct {
	CallerNumber   string   `json:"caller_number" validate:"required"`
	CallDuration   int      `json:"call_duration" validate:"required"`
	RecordingURL   string   `json:"recording_url,omitempty"`
	Transcript     string   `json:"transcript,omitempty"`
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	RiskScore      *float64 `json:"risk_score,omitempty"`
	Flagged        bool     `json:"flagged"`
	Notes          string   `json:"notes,omitempty"`
	TagIDs         []string `json:"tag_ids,omitempty"`
}

// UpdateCallLogRequest represents a request to update a call log
type UpdateCallLogRequest struct {
	CallerNumber   string   `json:"caller_number,omitempty"`
	CallDuration   int      `json:"call_duration,omitempty"`
	RecordingURL   string   `json:"recording_url,omitempty"`
	Transcript     string   `json:"transcript,omitempty"`
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	RiskScore      *float64 `json:"risk_score,omitempty"`
	Flagged        bool     `json:"flagged,omitempty"`
	Notes          string   `json:"notes,omitempty"`
	TagIDs         []string `json:"tag_ids,omitempty"`
}
