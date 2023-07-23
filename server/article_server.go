package server

import (
	"fmt"
	"mygin/core"
	"mygin/model"
	"mygin/util"
)

func GetListService() core.Response {
	data := []model.ArticleModel{
		{
			Id:      util.GenID(),
			Title:   "Java入门",
			Content: "<Java入门>",
		},
		{
			Id:      util.GenID(),
			Title:   "Golang入门",
			Content: "<Golang入门>",
		},
		{
			Id:      util.GenID(),
			Title:   "C++入门",
			Content: "<C++入门>",
		},
	}
	return core.Success(data)
}
func GetDetailService(id int64) core.Response {
	data := model.ArticleModel{
		Id:      id,
		Title:   "Java入门",
		Content: "<Java入门>",
	}
	return core.Success(data)
}

func ModifyDetailService(id int64, articleModel model.ArticleModel) core.Response {
	fmt.Println(articleModel)
	data := model.ArticleModel{
		Id:      id,
		Title:   articleModel.Title,
		Content: articleModel.Content,
	}

	return core.Success(data)
}

func CreateService(article model.ArticleModel) core.Response {
	article.Id = util.GenID()
	// 保存
	// *****
	return core.Success(article)
}

func DeleteService(id int64) core.Response {
	return core.Success(true)
}
