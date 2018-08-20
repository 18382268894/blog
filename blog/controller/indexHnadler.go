/**
*FileName: controller
*Create on 2018/8/20 下午5:47
*Create by mok
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context){
	c.HTML(http.StatusOK,"/views/index.html",nil)
}
