package model

import (
	"time"

	"github.com/luoluoluo/ws_api/global"
)

// Comment comment model
type Comment struct {
	ID         int    `json:"id" db:"id"`
	TaskID     int    `json:"task_id" db:"task_id"`
	UserID     int    `json:"user_id" db:"user_id"`
	UserName   string `json:"user_name" db:"user_name"`
	UserAvatar string `json:"user_avatar" db:"user_avatar"`
	Text       string `json:"text" db:"text"`
	Status     int    `json:"status" db:"status"`
	CreateTime string `json:"create_time" db:"create_time"`
}

// Add 新增comment
func (c *Comment) Add(taskID int, userID int, text string) (int, error) {
	res, err := global.DB.Exec(
		"INSERT INTO comment(task_id, user_id, text, status, create_time) VALUES(?,?,?,?,?)",
		taskID,
		userID,
		text,
		1,
		time.Now().Unix(),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
