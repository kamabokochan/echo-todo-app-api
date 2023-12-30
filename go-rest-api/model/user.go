package model

import "time"

// primaryKey = 主キー
// unique = 重複を許可しない
type User struct {
	ID        uint      `json:"id" grom:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 新しくユーザ生成した際の型
type UserResponse struct {
	ID    uint   `json:"id" grom:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
