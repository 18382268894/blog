/**
*FileName: blog
*Create on 2018/8/20 下午5:47
*Create by mok
*/

package main

import (
	"github.com/gin-gonic/gin"
	"blog/controller"
)

func main(){
	router := gin.Default()
	router.Static("/static/","./static")
	router.LoadHTMLGlob("views/*")

	//路由接口

	//1.首页
	router.GET("/",controller.IndexHandle)

	//2.文章发表页
	router.GET("/article/new",controller.NewArticle)
	router.POST("article/post_article",controller.PostAarticle)




	router.Run(":8080")
}
