package model

import "time"

// foreingnKey 他のテーブルの主キーを参照 一対多のリレーションを張る（複数のテーブル間で共通したカラム（列）を相互に関連づける）
// not null 空の値を許可しない
// OnDelete:CASCADE 参照先の行が削除された際の動作を指定するオプション。親テーブルの行が削除されたときに、子テーブルの関連する行も削除される

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreingnKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

// クライアントへ返すTask型
type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
