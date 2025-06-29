package models

import "time"

type Note struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     string    `gorm:"not null" json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CoverImage string    `json:"cover_image"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	IsPublic   bool      `gorm:"default:false" json:"is_public"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
