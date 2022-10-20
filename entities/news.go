package entities

import "time"

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"<-:created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
