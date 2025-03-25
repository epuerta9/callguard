package repository

import (
	"context"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/google/uuid"
)

// TagRepository provides access to the tag data store
type TagRepository struct {
	db *db.Queries
}

// NewTagRepository creates a new TagRepository
func NewTagRepository(db *db.Queries) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

// GetByID retrieves a tag by its ID
func (r *TagRepository) GetByID(ctx context.Context, id string) (*model.Tag, error) {
	// Convert id string to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbTag, err := r.db.GetTagByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return convertDBTagToTag(dbTag), nil
}

// List retrieves all tags
func (r *TagRepository) List(ctx context.Context) ([]*model.Tag, error) {
	// Placeholder implementation
	tags := []*model.Tag{
		{
			ID:        uuid.New().String(),
			Name:      "Suspicious",
			CreatedAt: "2023-06-01T10:00:00Z",
		},
		{
			ID:        uuid.New().String(),
			Name:      "Fraudulent",
			CreatedAt: "2023-06-01T10:05:00Z",
		},
		{
			ID:        uuid.New().String(),
			Name:      "Important",
			CreatedAt: "2023-06-01T10:10:00Z",
		},
	}
	return tags, nil
}

// Create creates a new tag
func (r *TagRepository) Create(ctx context.Context, req *model.CreateTagRequest) (*model.Tag, error) {
	// Placeholder implementation
	return &model.Tag{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: "2023-06-01T10:00:00Z",
	}, nil
}

// Update updates an existing tag
func (r *TagRepository) Update(ctx context.Context, id string, req *model.UpdateTagRequest) (*model.Tag, error) {
	// Placeholder implementation
	return &model.Tag{
		ID:        id,
		Name:      req.Name,
		CreatedAt: "2023-06-01T10:00:00Z",
	}, nil
}

// Delete deletes a tag
func (r *TagRepository) Delete(ctx context.Context, id string) error {
	// Placeholder implementation
	return nil
}

// GetTagsForCallLog retrieves all tags for a call log
func (r *TagRepository) GetTagsForCallLog(ctx context.Context, callLogID string) ([]*model.Tag, error) {
	// Placeholder implementation
	tags := []*model.Tag{
		{
			ID:        uuid.New().String(),
			Name:      "Suspicious",
			CreatedAt: "2023-06-01T10:00:00Z",
		},
		{
			ID:        uuid.New().String(),
			Name:      "Fraudulent",
			CreatedAt: "2023-06-01T10:05:00Z",
		},
	}
	return tags, nil
}

// convertDBTagToTag converts a db.Tag to a model.Tag
func convertDBTagToTag(dbTag db.Tag) *model.Tag {
	return &model.Tag{
		ID:        dbTag.ID.String(),
		Name:      dbTag.Name,
		CreatedAt: dbTag.CreatedAt.Time.Format(time.RFC3339),
	}
}
