package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/luoluoluo/ws_api/global"
)

// Task task model
type Task struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	UserName   string    `json:"user_name" db:"user_name"`
	UserAvatar string    `json:"user_avatar" db:"user_avatar"`
	Text       string    `json:"text" db:"text"`
	Images     string    `json:"images" db:"images"`
	Status     int       `json:"status" db:"status"`
	CreateTime int       `json:"create_time" db:"create_time"`
	Comments   []Comment `json:"comments"`
}

// TaskPaginator 分页
type TaskPaginator struct {
	Items   []Task `json:"items"`
	HasMore bool   `json:"has_more"`
}

// Tasks 获取最近size条task列表
func (t *Task) Tasks(mode int, userID int, lastID int, size int) (*TaskPaginator, error) {
	var tasks = []Task{}
	var hasMore = false
	var count int
	var taskCountSQL string
	var taskSQL string
	switch mode {
	case 0: // 全部的
		taskSQL = `SELECT t.*, u.name AS user_name, u.avatar AS user_avatar FROM task t 
			LEFT JOIN user u ON t.user_id=u.id 
			WHERE t.status=1`
		taskCountSQL = "SELECT COUNT(*) FROM task t WHERE t.status=1"
	case 1: // 我发布的
		taskSQL = `SELECT t.*, u.name AS user_name, u.avatar AS user_avatar FROM task t 
			LEFT JOIN user u ON t.user_id=u.id WHERE t.status=1  AND t.user_id=` + strconv.Itoa(userID)
		taskCountSQL = `SELECT COUNT(*) FROM task t 
			WHERE t.status=1 AND t.user_id=` + strconv.Itoa(userID)
	case 2: // 我评论的
		taskSQL = `SELECT t.*, u.name AS user_name, u.avatar AS user_avatar FROM task t 
			LEFT JOIN user u ON t.user_id=u.id 
			LEFT JOIN comment c ON t.id=c.task_id 
			WHERE t.status=1  AND c.user_id=` + strconv.Itoa(userID)
		taskCountSQL = `SELECT COUNT(*) FROM task t 
			LEFT JOIN comment c ON t.id = c.task_id 
			WHERE t.status=1 AND c.user_id=` + strconv.Itoa(userID)
	default:
		return &TaskPaginator{
			Items:   tasks,
			HasMore: hasMore,
		}, nil
	}

	if lastID != 0 {
		taskCountSQL += " AND t.id < " + strconv.Itoa(lastID)
	}

	err := global.DB.Get(&count, taskCountSQL)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &TaskPaginator{
			Items:   tasks,
			HasMore: hasMore,
		}, nil
	}
	if count > size {
		hasMore = true
	}

	if lastID != 0 {
		taskSQL += " AND t.id < " + strconv.Itoa(lastID)
	}
	taskSQL += " ORDER BY t.id DESC LIMIT " + strconv.Itoa(size)

	err = global.DB.Select(&tasks, taskSQL)

	if err != nil {
		return nil, err
	}

	var taskIDArr []string
	for _, task := range tasks {
		taskIDArr = append(taskIDArr, strconv.Itoa(task.ID))
	}
	taskIDStr := strings.Join(taskIDArr, ",")
	commentSQL := "SELECT c.*, u.name as user_name, u.avatar as user_avatar FROM comment c LEFT JOIN user u ON c.user_id=u.id WHERE c.status=1 AND c.task_id IN(" + taskIDStr + ")"
	comments := []Comment{}
	global.DB.Select(&comments, commentSQL)

	for i, task := range tasks {
		for _, comment := range comments {
			if task.ID == comment.TaskID {
				tasks[i].Comments = append(tasks[i].Comments, comment)
			}
		}
	}
	return &TaskPaginator{
		Items:   tasks,
		HasMore: hasMore,
	}, nil
}

// Add 新增task
func (t *Task) Add(userID int, text string) (int, error) {
	res, err := global.DB.Exec(
		"INSERT INTO task(user_id, text, status, create_time) VALUES(?,?,?,?)",
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
