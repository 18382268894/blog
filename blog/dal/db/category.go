/**
*FileName: db
*Create on 2018/8/20 上午10:39
*Create by mok
*/

package db

import (
	"blog/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)


func InsertCategory(category model.Category)(categoryid int,err error){
	sqlstr := `insert into category(category_name,category_no) values(?,?)`
	result,err := DB.Exec(sqlstr,category.CategoryName,category.CategoryNo)
	if err !=nil{
		err = fmt.Errorf("insert category failed,err:%v",err)
		return -1,err
	}
	lastid,err  := result.LastInsertId()
	if err !=nil{
		err = fmt.Errorf("get categoryid failed,err:%v",err)
		return -1,err
	}
	return int(lastid),nil
}

func GetCategoryList(categoryids []int)(categorylist []*model.Category,err error){
	sqlstr,args,err := sqlx.In(`select category_name,category_no from category where id in(?)`,categoryids)
	if err != nil{
		return nil,err
	}
	err = DB.Select(&categorylist,sqlstr,args...)
	return
}

func GetAllCategoryList()(categorylist []*model.Category,err error){
	sqlstr := `select id,category_name,category_no from category`
	err = DB.Select(&categorylist,sqlstr)
	if err != nil{
		return nil,err
	}
	return
}


func GetCategoryByID(categoryid int)(category model.Category,err error){
	if categoryid < 0 {
		err = fmt.Errorf("invalid parameter,categoryid:%d", categoryid)
		return
	}
	sqlstr := `select category_name,category_no from category where id=?`
	err = DB.Get(&category,sqlstr,categoryid)
	return
}

func GetRelativeArticleByID(categoryid int)(articlelist []*model.RelativeArticle,err error){
	if categoryid < 0 {
		err = fmt.Errorf("invalid parameter,categoryid:%d", categoryid)
		return
	}
	sqlstr := `select id,title from article where category_id=?`
	err = DB.Select(&articlelist,sqlstr,categoryid)
	return
}