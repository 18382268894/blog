/**
*FileName: db
*Create on 2018/8/20 上午11:33
*Create by mok
*/

package db

import (
	"testing"
	"blog/model"
	"fmt"
)

func TestInsertComment(t *testing.T) {
	var  comment  model.Comment
	comment.ArticleID = 2
	comment.Content = "哈哈哈，打不过我吧"
	comment.UserName = "特种兵"
	commentid,err := InsertComment(comment)
	if err != nil{
		t.Logf("insert comment failed,err:%v",err)
		return
	}
	fmt.Println(commentid)
}

func TestGetCommentList(t *testing.T) {
	commnetlist,err := GetCommentList(2,0,6)
	if err != nil{
		t.Logf("get commentlist failed,err:%v",err)
		return
	}
	for _,comment := range commnetlist{
		fmt.Println(comment.Content)
	}
}