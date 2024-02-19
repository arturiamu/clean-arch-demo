package models

type Article struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at" gorm:"index"`
}
