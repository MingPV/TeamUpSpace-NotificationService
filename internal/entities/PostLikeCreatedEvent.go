package entities

import "github.com/google/uuid"

type PostLikeCreatedEvent struct {
	PostId      int       `gorm:"primaryKey" json:"post_id"`
	PostOwnerId uuid.UUID `gorm:"type:uuid" json:"post_owner_id"`
	UserId      uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
}
