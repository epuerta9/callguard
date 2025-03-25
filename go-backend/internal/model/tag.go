package model

// Tag represents a tag for categorizing call logs
type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// CreateTagRequest represents a request to create a tag
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateTagRequest represents a request to update a tag
type UpdateTagRequest struct {
	Name string `json:"name" validate:"required"`
}
