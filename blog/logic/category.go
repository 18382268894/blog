/**
*FileName: logic
*Create on 2018/8/20 下午7:40
*Create by mok
*/

package logic

import (
	"blog/model"
	"blog/dal/db"
)

//从数据库中获取分类列表
func GetAllCategoryList()(categoryList []*model.Category,err error){
	categoryList,err = db.GetAllCategoryList()
	return
}