package response

import (
	"github.com/gin-gonic/gin"
	"golang-test/lanproxy-go-server/app/model"
	"net/http"
)

// 通用api响应
type TableResp struct {
	t *model.TableDataInfo
	c *gin.Context
}

// 返回一个成功的消息体
func BuildTable(c *gin.Context, total int, rows interface{}) *TableResp {
	msg := model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		t: &msg,
		c: c,
	}
	return &a
}

// 输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOk, resp.t)
	resp.c.Abort()
}
