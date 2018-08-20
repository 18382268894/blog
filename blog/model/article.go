/**
*FileName: model
*Create on 2018/8/14 上午11:31
*Create by mok
*/

package model

import "time"

type ArticleInfo struct {
	ID int64 `db:"id"`
	CategoryID int `db:"category_id"`
	Summary string `db:"summary"`
	Title string `db:"title"`
	ViewCount int64 `db:"view_count"`
	CommentCount int64 `db:"comment_count"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	UserName string `db:"username"`
}

type ArticleDetail struct {
	Content string `db:"content"`
	ArticleRecord
}


type ArticleRecord struct {
	ArticleInfo
	Category
}

