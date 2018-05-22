package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/model"
)

// TaskController task 控制器
type TaskController struct {
	Controller
}

type timelineReq struct {
	// LastID 最后一条记录的id
	LastID int `json:"last_id"`
	Size   int `json:"size"`
}

type addTaskReq struct {
	Text string `json:"text"`
}

// Timeline 时间线
func (tc *TaskController) Timeline(c *gin.Context) {
	req := &timelineReq{}
	c.BindJSON(req)
	if req.Size == 0 {
		req.Size = 10
	}
	task := &model.Task{}
	paginator, err := task.Tasks(req.LastID, req.Size)
	if err != nil {
		glog.Error(err)
		tc.resp(c, 500, gin.H{})
		return
	}
	tc.resp(c, 200, paginator)
	return
}

// Add 新增task
func (tc *TaskController) Add(c *gin.Context) {
	req := &addTaskReq{}
	c.BindJSON(req)
	if req.Text == "" {
		tc.resp(c, 400, gin.H{})
		return
	}
	task := &model.Task{}
	user := c.MustGet("user").(*model.User)
	id, err := task.Add(user.ID, req.Text)
	if err != nil {
		glog.Error(err)
		tc.resp(c, 500, gin.H{})
		return
	}
	tc.resp(c, 201, gin.H{
		"id": id,
	})
}
