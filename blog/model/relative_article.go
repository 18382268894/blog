/**
*FileName: model
*Create on 2018/8/20 上午10:19
*Create by mok
 */

package model

type RelativeArticle struct {
	Articleid int64  `db:"id"`
	Title     string `db:"title"`
}
