package models

import "time"

type Profile struct {
	ID       string    `gorm:"type:uuid;primaryKey" json:"id"`
	FullName string    `json:"full_name"`
	Bio      string    `json:"bio"`
	CoverUrl string    `json:"cover_url"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
}
