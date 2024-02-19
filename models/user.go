package models

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at" gorm:"index"`
}
