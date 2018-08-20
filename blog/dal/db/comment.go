/**
*FileName: db
*Create on 2018/8/20 上午11:32
*Create by mok
*/

package db

import (
	"blog/model"
	"fmt"
)

func InsertComment(comment model.Comment)(commentid int64,err error){
	if comment.ArticleID < 0{
		err = fmt.Errorf("invalid parameter，articleid:%d",comment.ArticleID)
		return -1,err
	}
	tx,err := DB.Beginx()
	defer func() {
		if err != nil{
			tx.Rollback()
		}
	}()

	sqlstr := `insert into comment(article_id,content,username) values(?,?,?)`
	result,err := tx.Exec(sqlstr,comment.ArticleID,comment.Content,comment.UserName)
	if err != nil{
		return -1,err
	}
	commentid,err = result.LastInsertId()
	if err != nil{
		return -1,err
	}
	sqlstr = `update article set comment_count=comment_count+1 where id=?`
	_,err = tx.Exec(sqlstr,comment.ArticleID)
	if err != nil{
		return -1,err
	}
	err = tx.Commit()
	if err != nil{
		return -1,err
	}
	return
}

func GetCommentList(articleid int64,pageNum int,pageSize int)(commentlist []*model.Comment,err error){
	if articleid < 0{
		err = fmt.Errorf("invalid parameter,articleid:%d",articleid)
		return
	}
	if pageNum < 0 || pageSize < 0{
		err = fmt.Errorf("invalid parameter,pageNum:%d,pageSize:%d",pageNum,pageSize)
		return
	}
	sqlstr := `select article_id,username,content,create_time from comment where article_id=? and status=1 order by create_time asc limit ?,?`
	err = DB.Select(&commentlist,sqlstr,articleid,pageNum,pageSize)
	return
}