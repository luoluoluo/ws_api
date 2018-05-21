package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/library"
)

// TaskController task 控制器
type TaskController struct {
	c Controller
}

type listReq struct {
	// LastID 最后一条记录的id
	LastID int `json:"last_id"`
}

// List 时间线
func (t *TaskController) List(c *gin.Context) {
	req := &listReq{}
	c.BindJSON(req)
	db := c.MustGet("db").(*library.DB)

	var tasks []map[string]string
	var err error
	if req.LastID == 0 {
		tasks, err = db.Select("SELECT * FROM task ORDER BY id DESC LIMIT 10")
	} else {
		tasks, err = db.Select("SELECT * FROM task WHERE id <? ORDER BY id DESC LIMIT 10", req.LastID)
	}
	if err != nil {
		return
	}

	var taskIds []int
	for i, task := range tasks {
		taskIds[i] = library.ParseInt(task["id"])
	}

}

// Info 详情
func (t *TaskController) Info(c *gin.Context) {
}

// Post 发布
func (t *TaskController) Post(c *gin.Context) {
}
