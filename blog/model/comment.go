/**
*FileName: model
*Create on 2018/8/14 上午11:46
*Create by mok
*/

package model

import "time"

type Comment struct {
	ID int64 `db:"id"`
	UserName string`db:"username"`
	Content string `db:"content"`
	CteateTime time.Time `db:"create_time"`
	Status int `db:"status"`
	ArticleID int64 `db:"article_id"`
}