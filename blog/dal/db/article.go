/**
*FileName: article
*Create on 2018/8/14 下午12:16
*Create by mok
*/

package db

import (
	"blog/model"
	"fmt"
)

//插入文章
func InsertArticle(article *model.ArticleDetail)(articleid int64,err error){
	if article == nil{
		return -1,fmt.Errorf("invalid parameter")
	}
	sqlstr := `insert into article(
					category_id,content,title,view_count,comment_count,username,summary
			   )values(?,?,?,?,?,?,?)
			  `
	relsult,err := DB.Exec(sqlstr,article.ArticleInfo.CategoryID,article.Content,article.Title,article.ViewCount,article.CommentCount,article.UserName,article.Summary)
	if err != nil{
		return -1,err
	}
	articleid,err =relsult.LastInsertId()
	return
}

//获取所有文章列表
func GetArticleList(pageNum int,pageSize int)(articleList []*model.ArticleInfo,err error){
	if pageNum < 0 || pageSize < 0{
		return nil,fmt.Errorf("invalied pagenum or pagesize")
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList,sqlstr,pageNum,pageSize)
	if err != nil{
		return nil,err
	}
	return
}

//获取文章信息
func GetArticleDetail(articleid int64)(ariticleDetail *model.ArticleDetail,err error){
	if articleid < 0{
		err = fmt.Errorf("invalid parameter,articleid:%d",articleid)
		return
	}
	ariticleDetail = new(model.ArticleDetail)
	sqlstr := `
				select id,category_id,content,title,view_count,comment_count,username,summary,create_time
				from article
				where id=? and status=1`
	err  = DB.Get(ariticleDetail,sqlstr,articleid)
	if err != nil{
		return nil,err
	}
	return
}

//获取相关文章
func GetRelativeArticle(articleid int64)(articlelist []*model.RelativeArticle,err error){
	if articleid < 0{
		err = fmt.Errorf("invalid parameter,articleid:%d",articleid)
		return
	}
	var categoryid int
	sqlstr := `select category_id from article where id=? and status=1`
	err = DB.Get(&categoryid,sqlstr,articleid)
	if err != nil{
		return nil,err
	}
	sqlstr = `select id,title from article where category_id=? and status=1 and id!=? order by view_count desc limit 10`
	err = DB.Select(&articlelist,sqlstr,categoryid,articleid)
	if err != nil{
		return nil,err
	}
	return
}