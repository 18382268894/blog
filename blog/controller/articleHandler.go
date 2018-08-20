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
	"strconv"
)

func NewArticle(c *gin.Context){
	categoryList,err := logic.GetAllCategoryList()
	if err != nil{
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.HTML(http.StatusOK,"views/newarticle.html",categoryList)
}

func PostArticle(c *gin.Context){
	username := c.PostForm("username")
	content := c.PostForm("content")
	title := c.PostForm("title")
	categoryidstr := c.PostForm("categoryid")
	categoryid, err := strconv.Atoi(categoryidstr)
	if err != nil{
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	_, err = logic.InserArticle(content,username,title,categoryid)
	if err !=nil{
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently,"/")
}