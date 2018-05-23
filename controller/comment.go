package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/model"
)

// CommentController comment控制器
type CommentController struct {
	Controller
}

type addCommentReq struct {
	Text   string `json:"text"`
	TaskID int    `json:"task_id"`
}

// Add 新增评论
func (cc *CommentController) Add(c *gin.Context) {
	req := &addCommentReq{}
	c.BindJSON(req)
	comment := &model.Comment{}
	user := c.MustGet("user").(*model.User)
	id, err := comment.Add(req.TaskID, user.ID, req.Text)
	if err != nil {
		glog.Error(err)
		cc.resp(c, 500, gin.H{})
		return
	}
	cc.resp(c, 201, gin.H{
		"id": id,
	})
}
