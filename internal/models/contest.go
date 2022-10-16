package models

import "time"

type Contest struct {
	ContestId   int32     `json:"contestID"`
	Title       string    `json:"title"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Defunct     string    `json:"defunct"`
	Description string    `json:"description"`
	Private     uint8     `json:"private"`
	Langmask    int       `json:"langmask"`
	Password    string    `json:"password"`
	UserId      int32     `json:"user_id"`
}

func (Contest) TableName() string {

	return "contest"
}

type AddContest struct {
	Cid         int32    `json:"cid"`
	Title       string   `json:"title"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	Description string   `json:"description"`
	ProIDS      string   `json:"proIDS"`
	Role        int32    `json:"role"`
	LimitUser   []string `json:"limitUser"`
}
