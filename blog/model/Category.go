/**
*FileName: model
*Create on 2018/8/14 上午11:43
*Create by mok
*/

package model

type Category struct {
	CategoryID int `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo  int  `db:"category_no"`
}