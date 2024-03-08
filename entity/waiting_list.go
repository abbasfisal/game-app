package entity

type WaitingMember struct {
	UserID    uint     `json:"user_id"`
	Timestamp int64    `json:"timestamp"`
	Category  Category `json:"category"`
}
