/**
*FileName: db
*Create on 2018/8/14 下午12:35
*Create by mok
*/

package db

import (
	"testing"
	"fmt"
	"blog/model"
)




func TestInsertArticle(t *testing.T) {
	article := new(model.ArticleDetail)
	article.UserName = "Mok"
	article.Title = "这是一篇测试文章"
	article.Summary = "这是一篇测试文章"
	article.Content = "这是一篇测试文章"
	article.ViewCount = 0
	article.CommentCount = 0
	article.ArticleInfo.CategoryID = 1
	articleid,err := InsertArticle(article)
	if err != nil{
		t.Logf("insert article falied,err:%v",err)
		return
	}
	t.Logf("insert artilce success,artilce id :%d",articleid)
}

func TestGetArticleList(t *testing.T) {
	articles,err := GetArticleList(1,15)
	if err != nil{
		t.Logf("Get article list falied,err:%v",err)
		return
	}
	fmt.Println(len(articles))
}

func TestGetArticleDetail(t *testing.T) {
	var articleid int64= 5
	articledetail,err := GetArticleDetail(articleid)
	if err != nil{
		t.Logf("Get article detail failed,err:%v",err)
		return
	}
	fmt.Println(articledetail)
}

func TestGetRelativeArticle(t *testing.T) {
	var articleid int64= 14
	articlelist,err := GetRelativeArticle(articleid)
	if err != nil{
		t.Logf("Get article detail failed,err:%v",err)
		return
	}
	fmt.Println(len(articlelist))
}