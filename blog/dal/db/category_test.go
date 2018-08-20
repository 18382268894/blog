/**
*FileName: db
*Create on 2018/8/20 上午10:48
*Create by mok
*/

package db

import (
	"testing"
	"blog/model"
	"fmt"
)


func TestInsertCategory(t *testing.T) {
	var category model.Category
	category.CategoryNo = 7
	category.CategoryName = "数据库管理"
	categoryid ,err := InsertCategory(category)
	if err != nil{
		t.Logf("insert category falied,err:%v",err)
		return
	}
	fmt.Println(categoryid)
}

func TestGetCategoryList(t *testing.T) {
	var categoryids = []int{2,3,4,5}
	categorylist,err := GetCategoryList(categoryids)
	if err != nil{
		t.Logf("Get category falied,err:%v",err)
		return
	}
	for _,category := range categorylist{
		fmt.Println(category.CategoryName)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categorylist ,err := GetAllCategoryList()
	if err != nil{
		t.Logf("Get category falied,err:%v",err)
		return
	}
	for _,category := range categorylist{
		fmt.Println(category.CategoryName)
	}
}

func TestGetCategoryByID(t *testing.T) {
	category,err := GetCategoryByID(3)
	if err != nil {
		t.Logf("Get category failed,err:%v",err)
		return
	}
	fmt.Println(category)
}

func TestGetRelativeArticleByID(t *testing.T) {
	articlelist,err := GetRelativeArticleByID(3)
	if err != nil {
		t.Logf("Get relative article failed,err:%v",err)
		return
	}
	for _,article:= range articlelist{
		fmt.Println(article.Title)
	}
}