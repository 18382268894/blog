/**
*FileName: controller
*Create on 2018/8/20 下午5:47
*Create by mok
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/logic"
)

func IndexHandle(c *gin.Context){
	articleRecordList,err := logic.GetArticleRecordList()
	if err != nil{
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
	}
	c.HTML(http.StatusOK,"views/index.html",articleRecordList)
}
