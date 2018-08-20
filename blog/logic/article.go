/**
*FileName: logic
*Create on 2018/8/20 下午7:40
*Create by mok
*/

package logic

import (
	"blog/model"
	"blog/dal/db"
	"fmt"
	"math"
)

func InserArticle(content,username,title string,categoryid int)(articleid int64,err error){
	if categoryid < 0{
		err = fmt.Errorf("invalid parameter,categoryid:%d",categoryid)
		return -1,err
	}
	length := int(math.Min(float64(len([]rune(content))),128.00))
	summary := string([]rune(content)[:length])
	articleDetail := &model.ArticleDetail{
		Content:content,
		ArticleRecord:model.ArticleRecord{
			ArticleInfo:model.ArticleInfo{
				UserName:username,
				Title:title,
				CategoryID:categoryid,
				Summary:summary,
			},
		},
	}
	articleid,err = db.InsertArticle(articleDetail)
	return
}


func GetArticleRecordList()(articleRecordList []*model.ArticleRecord,err error){
	articleInfoList,err := db.GetArticleList(0,10)
	if err != nil{
		return nil,err
	}
	var category model.Category
	for _,articleInfo := range articleInfoList{
		category ,err = db.GetCategoryByID(articleInfo.CategoryID)
		if err != nil{
			return nil,err
		}
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*articleInfo,
			Category:category,
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}
